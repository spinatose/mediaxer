package fileops

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"testing"
)

const tempFolder string = "../tmp"

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

func TestGetFileThumbnails(t *testing.T) {
	err := createTempFilesAndCleanup(t, 4)

	if err != nil {
		t.Errorf("TestGetFileThumbnails unable to setup temp folder with test files [%s] error: %s\n", tempFolder, err)
	}
}

func createTempFilesAndCleanup(t *testing.T, quantity int) error {
	var err error = nil 

	for i := 0; i < quantity; i++ {
		for k := 0; k < quantity; k++ {
			fl := path.Join(tempFolder, "tmp" + strconv.Itoa(i), "temp" + strconv.Itoa(k) + ".txt")
			err = os.MkdirAll(filepath.Dir(fl), 0770); 
			
			if err != nil {
				break 
			}

			err = os.WriteFile(fl, []byte(strconv.Itoa(i+k)), os.ModePerm)
			
			if err != nil {
				break
			}
		}

		if err != nil {
			break 
		}
	}

	t.Cleanup(func() {
		os.RemoveAll(tempFolder)
	})

	return err
}