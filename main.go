package main

import (
	"fmt"
	"os"
	"time"
	"spinatose.com/mediaxer/fileops" 
)

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
		} else {
			folder := args[0]
			fmt.Printf("Folder value provided from command line: %s\n", folder);
			fmt.Printf("Folder value provided is directory and is accessible? %t\n", fileops.ValidMachineFolder(folder)) 
		}
	}

	fmt.Println()
	fmt.Println(time.Now().Format("Mon Jan 2 15:04:05 MST 2006"))
}
