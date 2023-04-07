package fileops

import (
	"fmt"
	"testing"
)

func TestValidMachineFolder(t *testing.T) {
	// ARRANGE
	folder := "not a valid folder at all"
	
	//ACT 
	// should return that folder is invalid- false
	validFolder, err := ValidMachineFolder(folder)

	// ASSERT
	if !validFolder && err != nil {
		fmt.Printf("TestValidMachineFolder SUCCESS for invalid folder [%s] - error: %s\n", folder, err.Error())
	} else {
		t.Errorf("TestValidMachineFolder failed for checking folder value [%s] - error: %s\n", folder, err.Error())
	}

	// ARRANGE
	folder = "."

	// ACT
	// should return that folder is valid- true
	validFolder, err = ValidMachineFolder(folder)

	// ASSERT
	if validFolder && err == nil {
		fmt.Printf("TestValidMachineFolder SUCCESS for valid folder [%s]\n", folder)
	} else {
		t.Errorf("TestValidMachineFolder failed for checking folder value [%s] - error: %s\n", folder, err.Error())
	}

	// ARRANGE
	folder = "fileops.go"

	// ACT
	// should return that folder is invalid- false, since folder is file not a folder
	validFolder, err = ValidMachineFolder(folder)

	// ASSERT
	if !validFolder && err != nil {
		fmt.Printf("TestValidMachineFolder SUCCESS for invalid folder [%s] - error: %s\n", folder, err.Error())
	} else {
		t.Errorf("TestValidMachineFolder failed for checking folder value [%s]\n", folder)
	}
}
