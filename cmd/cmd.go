package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	configuration "spinatose.com/mediaxer/config"
	"spinatose.com/mediaxer/fileops"
	"spinatose.com/mediaxer/logging"
)

// Configuration file
const (
	configFile    string = "config.json"
	processFolder string = "./process"
)

var rootCmd = &cobra.Command{
	Use:   "mediaxer",
	Short: "MediAxer",
	RunE:  runApp,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	fs := rootCmd.PersistentFlags()
	fs.StringP("source", "s", "", "specify source folder")
	fs.StringP("dest", "d", "", "specify destination folder")
	//fs.BoolP("debug", "d", false, "enable debug profiling")

	if err := viper.BindPFlags(fs); err != nil {
		panic(err)
	}
}

func runApp(cmd *cobra.Command, args []string) error {
	fmt.Println("      Welcome to mediAxer - file organizer!")
	fmt.Println("¡Bienvenido a mediAxer - organizador de archivos!")
	fmt.Println()

	// Check for local processed folder, if not exists- create
	err := ensureLocalProcessFolderExists()
	if err != nil {
		fmt.Printf("Unable to ensure/create local 'processed' folder- error: %s", err.Error())
		return err
	}

	// Get any passed override arguments
	sourceFolder := viper.GetString("source")
	destFolder := viper.GetString("dest")

	// Get config or create default and load it.
	config, err := getAppConfig()
	if err != nil {
		return err
	}

	logger := logging.NewLogger(config.Logger.Outputs)
	logger.Debug("GOT HERE BOY!!")
	logger.Info("trying with the info")


	err = resolveAppArgsConfig(config, sourceFolder, destFolder)
	if err != nil {
		fmt.Printf("Unable to validate/resolve settings loaded for applicaion- error: %s", err.Error())
		return err
	}

	// TODO: Use logging to have this only show in debug loglevel
	fmt.Print("configuration loaded...\n")
	fmt.Println(config.ToString())

	fmt.Println()
	fmt.Println(time.Now().Format("Mon Jan 2 15:04:05 MST 2006"))

	return nil
}

func ensureLocalProcessFolderExists() error {
	validFolder, _ := fileops.ValidMachineFolder(processFolder)

	if !validFolder {
		if err := os.Mkdir(processFolder, os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}

func getAppConfig() (*configuration.Config, error) {
	config := configuration.NewConfig()

	// Check for config file existence
	if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
		// Config file doesn't exist, create one with default values.
		err = configuration.CreateConfigFile(config, configFile)

		if err != nil {
			fmt.Printf("Error creating default configuration file [%s]- error: %s\n", configFile, err.Error())
			return nil, err
		}
	} else {
		config, err = configuration.LoadConfigFromJsonFile(configFile)

		if err != nil {
			fmt.Printf("Error attempting to load configuration file [%s]- error: %s\n", configFile, err.Error())
			return nil, err
		}
	}

	return config, nil
}

func isArgProvided(name string) bool {
	name = strings.ToUpper(name)
	found := false
	args := os.Args[1:]

	for _, arg := range args {
		if strings.Contains(strings.ToUpper(arg), name) {
			found = true
		}
	}

	return found
}

func resolveAppArgsConfig(config *configuration.Config, sourceFolder string, destFolder string) error {
	// Override config settings with passed in arguments
	if isArgProvided("--source") || isArgProvided("-s") {
		config.SourceFolder = sourceFolder
	}

	if isArgProvided("--dest") || isArgProvided("-d") {
		config.DestinationFolder = destFolder
	}

	if config.SourceFolder == "" {
		return errors.New("an accessible, valid source folder must be supplied--  type '-help' for usage")
	} else {
		validFolder, err := fileops.ValidMachineFolder(config.SourceFolder)

		if validFolder {
			fmt.Printf("SourceFolder [%s] is a valid folder\n", config.SourceFolder)
		} else {
			return fmt.Errorf("sourcefolder [%s] not a valid folder- error: %s", config.SourceFolder, err.Error())
		}
	}

	if config.DestinationFolder == "" {
		config.DestinationFolder = config.SourceFolder
	} else {
		validFolder, err := fileops.ValidMachineFolder(config.DestinationFolder)

		if validFolder {
			fmt.Printf("DestinationFolder [%s] is a valid folder\n", config.DestinationFolder)
		} else {
			return fmt.Errorf("destinationfolder [%s] not a valid folder- error: %s", config.DestinationFolder, err.Error())
		}
	}

	return nil
}
