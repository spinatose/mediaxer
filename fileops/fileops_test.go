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

	// RECURSIVE WITH SUBFOLDERS
	// ACT
	thumbs, err := GetFileThumbnails(tempFolder, true, nil) 

	// ASSERT
	if thumbs == nil || err != nil{
		t.Errorf("TestGetFileThumbnails(recursive) failed to retrieve thumbnails from folder [%s]- error: %s", tempFolder, err)
	}

	lent := len(thumbs.Items)
	if lent != 20 {
		t.Errorf("TestGetFileThumbnails(recursive) failed to retrieve all thumbnails from folder [%s] found [%v]- error: %s", tempFolder, lent, err)
	}

	if thumbs.Items[0].Name != "temp0.txt" {
		t.Errorf("TestGetFileThumbnails(recursive) failed to retrieve name on thumbnails: expected [temp0.txt], found [%s]", thumbs.Items[0].Name)
	}

	if thumbs.Items[0].CreatedDate.Day() != time.Now().Day() || 
		thumbs.Items[0].CreatedDate.Hour() != time.Now().Hour() {
		t.Errorf("TestGetFileThumbnails(recursive) failed to retrieve correct created date on thumbnails: createdDate [%v]", thumbs.Items[0].CreatedDate)
	}

	if thumbs.Items[15].OriginPath != path.Join(tempFolder, "tmp2") {
		t.Errorf("TestGetFileThumbnails(recursive) failed to retrieve correct origin path on thumbnails: originPath [%s]", thumbs.Items[15].OriginPath)
	}

	// NON-RECURSIVE
	// ACT
	thumbs2, err := GetFileThumbnails(tempFolder, false, nil) 

	// ASSERT
	if thumbs2 == nil || err != nil{
		t.Errorf("TestGetFileThumbnails(non-recursive) failed to retrieve thumbnails from folder [%s]- error: %s", tempFolder, err)
	}

	lent2 := len(thumbs2.Items)
	if lent2 != 4 {
		t.Errorf("TestGetFileThumbnails(non-recursive) failed to retrieve all thumbnails from folder [%s] found [%v]- error: %s", tempFolder, lent, err)
	}

	if thumbs2.Items[0].Name != "temp0.txt" {
		t.Errorf("TestGetFileThumbnails(non-recursive) failed to retrieve name on thumbnails: expected [temp0.txt], found [%s]", thumbs2.Items[0].Name)
	}

	if thumbs2.Items[0].CreatedDate.Day() != time.Now().Day() || 
		thumbs2.Items[0].CreatedDate.Hour() != time.Now().Hour() {
		t.Errorf("TestGetFileThumbnails(non-recursive) failed to retrieve correct created date on thumbnails: createdDate [%v]", thumbs2.Items[0].CreatedDate)
	}

	if thumbs2.Items[3].OriginPath != tempFolder {
		t.Errorf("TestGetFileThumbnails(non-recursive) failed to retrieve correct origin path on thumbnails: originPath [%s]", thumbs2.Items[3].OriginPath)
	}
}

func TestGetCreatedDays(t *testing.T) {
	// ARRANGE
	err := createTempFilesAndCleanup(t, 4)
	thumbs, err := GetFileThumbnails(tempFolder, false, nil)

	if (err != nil) {
		t.Errorf("TestGetCreatedDays unable to prepare test assets with folder [%s]- error: %s", tempFolder, err)
	}

	// ACT
	days := thumbs.GetCreatedDays()

	// ASSERT
	if (len(days) != 1) {
		t.Errorf("TestGetCreatedDays failed to retrieve days from thumbs representing folder [%s], expected 1 day- actual [%v]- error: %s", tempFolder, len(days), err)
	}

	if (days[0] != time.Now().Format(time.DateOnly)) {
		t.Errorf("TestGetCreatedDays failed to retrieve days from thumbs representing folder [%s], wrong value for day expected [%s]- actual [%s]", tempFolder, time.Now().Format(time.DateOnly), days[0])
	}
}	

func createTempFilesAndCleanup(t *testing.T, quantity int) error {
	var err error = nil 

	err = os.Mkdir(tempFolder, 0770)

	for i := 0; i < quantity; i++ {
		// root file creation
		fl := path.Join(tempFolder, "temp" + strconv.Itoa(i) + ".txt")
		err = os.WriteFile(fl, []byte(strconv.Itoa(i)), os.ModePerm)
			
		if err != nil {
			break
		}
		// subfolder files
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