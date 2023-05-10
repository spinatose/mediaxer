package fileops

import (
	"io/ioutil"
	"os"
	"time"
)

func ValidMachineFolder(folder string) (bool, error) {
	_, err := os.ReadDir(folder)

	if err != nil {
		return false, err
	}

	return true, nil
}

type Thumbnail struct {
	CreatedDate time.Time
	Name        string
	OriginPath   string 
	Size        int64
}

func GetFileThumbnails(path string, thumbs []Thumbnail) ([]Thumbnail, error) {
	files, err := ioutil.ReadDir(path)
	
	// init return slice if not created
	if thumbs == nil {
		thumbs = []Thumbnail{}
	}

	if err != nil {
		return nil, err
	}

	for _, file := range files {
		filepath := path + "/" + file.Name()
		if file.IsDir() {
			thumbs, err = GetFileThumbnails(filepath, thumbs)
			if (err != nil) {
				return nil, err
			}
		} else {
			//if strings.Contains(file.Name(), ".txt") {
			thumbs = append(thumbs, Thumbnail{
				CreatedDate: file.ModTime(),
				Name: file.Name(),
				OriginPath: path,
				Size: file.Size(),
			})
			//}
		}
	}

	return thumbs, err
}