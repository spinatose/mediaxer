package fileops

import (
	"fmt"
	"testing"
)

func TestValidMachineFolder(t *testing.T) {
	folder := "not a valid folder at all"
	// should return that folder is invalid- false
	validFolder, err := ValidMachineFolder(folder)

	if !validFolder && err != nil {
		fmt.Printf("TestValidMachineFolder SUCCESS for invalid folder [%s] - error: %s\n", folder, err.Error())
	} else {
		t.Errorf("TestValidMachineFolder failed for checking folder value [%s] - error: %s\n", folder, err.Error())
	}

	folder = "."
	// should return that folder is valid- true
	validFolder, err = ValidMachineFolder(folder)

	if validFolder && err == nil {
		fmt.Printf("TestValidMachineFolder SUCCESS for valid folder [%s]\n", folder)
	} else {
		t.Errorf("TestValidMachineFolder failed for checking folder value [%s] - error: %s\n", folder, err.Error())
	}

	folder = "fileops.go"
	// should return that folder is invalid- false, since folder is file not a folder
	validFolder, err = ValidMachineFolder(folder)

	if !validFolder && err != nil {
		fmt.Printf("TestValidMachineFolder SUCCESS for invalid folder [%s] - error: %s\n", folder, err.Error())
	} else {
		t.Errorf("TestValidMachineFolder failed for checking folder value [%s]\n", folder)
	}
}
