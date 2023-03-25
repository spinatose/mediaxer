package main

import (
	"errors"
	"fmt"
	"os"
	"time"
	"spinatose.com/mediaxer/config"
	"spinatose.com/mediaxer/fileops" 
)

// TODO: eventually add command line args to filter certain files
// TODO: allow running from config file instead of command line- or allow override 
// TODO: provide flag type sets for filters like "all photos", "all videos", "all media", "all text files", etc...
// TODO: allow for argument for source and target directories (will have to check both for valid folders)
// TODO: possibly allow to upload to cloud
// TODO: possibly create file streaming server to accept incoming "files" stream over network and save to target
// TODO: create GUI interface

// Configuration file
const configFile string = "config.json"

func main() {
	fmt.Println("      Welcome to mediAxer - file organizer!")
	fmt.Println("¡Bienvenido a mediAxer - organizador de archivos!")
	fmt.Println() ;

	// arg at index 0 is the executable name
	args := os.Args[1:]
	
	if len(args) == 1 && args[0] == "-help" {
		fmt.Println("Usage: mediaxer <folder>")
		fmt.Println("Example: mediaxer '/users/bob/tmp/'")
	} else {
		if len(args) != 1 {
			fmt.Println("An accessible, valid folder must be supplied--  type '-help' for usage")
			return 
		} else {
			folder := args[0]
			fmt.Printf("Folder value provided from command line: %s\n", folder);
			fmt.Printf("Folder value provided is directory and is accessible? %t\n", fileops.ValidMachineFolder(folder)) 
		}
	}

	config := configuration.NewConfig()

	// Check for config file existence
	if _, err := os.Stat(configFile); errors.Is(err, os.ErrNotExist) {
		// Config file doesn't exist, create one with default values.
		err = configuration.CreateConfigFile(config, configFile)

		if err != nil {
			fmt.Printf("Error creating default configuration file [%s]- error: %s\n", configFile, err.Error())
			return 
		}
	} else {
		config, err = configuration.LoadConfigFromJsonFile(configFile)

		if err != nil {
			fmt.Printf("Error attempting to load configuration file [%s]- error: %s\n", configFile, err.Error())
			return 
		}
	}

	fmt.Print("configuration loaded...\n")
	fmt.Println(config.ToString())

	fmt.Println()
	fmt.Println(time.Now().Format("Mon Jan 2 15:04:05 MST 2006"))
}
