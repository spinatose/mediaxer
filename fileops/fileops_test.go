package fileops

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"testing"
	"time"
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
	// ARRANGE
	err := createTempFilesAndCleanup(t, 4)

	if err != nil {
		t.Errorf("TestGetFileThumbnails unable to setup temp folder with test files [%s]- error: %s\n", tempFolder, err)
	}

	// ACT
	thumbs, err := GetFileThumbnails(tempFolder, nil) 

	// ASSERT
	if thumbs == nil || err != nil{
		t.Errorf("TestGetFileThumnails failed to retrieve thumbnails from folder [%s]- error: %s", tempFolder, err)
	}

	lent := len(thumbs)
	if lent != 16 {
		t.Errorf("TestGetFileThumnails failed to retrieve all thumbnails from folder [%s] found [%v]- error: %s", tempFolder, lent, err)
	}

	if thumbs[0].Name != "temp0.txt" {
		t.Errorf("TestGetFileThumbnails failed to retrieve name on thumbnails: expected [temp0.txt], found [%s]", thumbs[0].Name)
	}

	if thumbs[0].CreatedDate.Day() != time.Now().Day() || 
		thumbs[0].CreatedDate.Hour() != time.Now().Hour() {
		t.Errorf("TestGetFileThumbnails failed to retrieve correct created date on thumbnails: createdDate [%v]", thumbs[0].CreatedDate)
	}

	if thumbs[15].OriginPath != path.Join(tempFolder, "tmp3") {
		t.Errorf("TestGetFileThumbnails failed to retrieve correct origin path on thumbnails: originPath [%s]", thumbs[15].OriginPath)
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