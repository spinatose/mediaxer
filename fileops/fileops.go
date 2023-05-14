package fileops

import (
	"io/ioutil"
	"os"
)

func ValidMachineFolder(folder string) (bool, error) {
	_, err := os.ReadDir(folder)

	if err != nil {
		return false, err
	}

	return true, nil
}

func GetFileThumbnails(path string, recursive bool, thumbs *Thumbnails) (*Thumbnails, error) {
	files, err := ioutil.ReadDir(path)
	
	// init return slice if not created
	if thumbs == nil {
		thumbs = NewThumbnails()
	}

	if err != nil {
		return nil, err
	}

	for _, file := range files {
		filepath := path + "/" + file.Name()
		if file.IsDir() {
			// if not recursive then skip subfolders
			if (!recursive) {
				continue
			}

			thumbs, err = GetFileThumbnails(filepath, recursive, thumbs)
			if (err != nil) {
				return nil, err
			}
		} else {
			//if strings.Contains(file.Name(), ".txt") {
			thumbs.AddItem(Thumbnail{
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