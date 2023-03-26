package cmd

import (
	"errors"
	"fmt"
	"os"
	"time"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"spinatose.com/mediaxer/config"
	"spinatose.com/mediaxer/fileops" 
)

// Configuration file
const configFile string = "config.json"

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
	fs.StringP("source", "s", "~/tmp", "specify source folder")
	//fs.BoolP("debug", "d", false, "enable debug profiling")

	if err := viper.BindPFlags(fs); err != nil {
		panic(err)
	}
}

func runApp(cmd *cobra.Command, args []string) error {
	fmt.Println("      Welcome to mediAxer - file organizer!")
	fmt.Println("¡Bienvenido a mediAxer - organizador de archivos!")
	fmt.Println() ;
	fmt.Printf("(-_-) received command line arguments: %s\n", args)
	// arg at index 0 is the executable name
	// args = args[1:]
	folder := viper.GetString("source")

	// if len(args) == 1 && args[0] == "-help" {
	// 	fmt.Println("Usage: mediaxer <folder>")
	// 	fmt.Println("Example: mediaxer '/users/bob/tmp/'")
	// } else {
		if folder == "" {
			fmt.Println("An accessible, valid folder must be supplied--  type '-help' for usage")
			return nil
		} else {
			//folder := args[0]
			validFolder, err := fileops.ValidMachineFolder(folder)

			if validFolder {
				fmt.Printf("Folder [%s] is a valid folder\n", folder)
			} else {
				fmt.Printf("Folder [%s] not a valid folder- error: %s\n", folder, err.Error()) 
			} 
		}
	//}

	config := configuration.NewConfig()

	// Check for config file existence
	if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
		// Config file doesn't exist, create one with default values.
		err = configuration.CreateConfigFile(config, configFile)

		if err != nil {
			fmt.Printf("Error creating default configuration file [%s]- error: %s\n", configFile, err.Error())
			return err
		}
	} else {
		config, err = configuration.LoadConfigFromJsonFile(configFile)

		if err != nil {
			fmt.Printf("Error attempting to load configuration file [%s]- error: %s\n", configFile, err.Error())
			return err
		}
	}

	fmt.Print("configuration loaded...\n")
	fmt.Println(config.ToString())

	fmt.Println()
	fmt.Println(time.Now().Format("Mon Jan 2 15:04:05 MST 2006"))

	return nil
}