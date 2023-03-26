package fileops

import (
	"os"
)

func ValidMachineFolder(folder string) (bool, error) {
	// TODO: check to see if is a valid directory
	// TODO: check to see if accessible
	// TODO: check to see if writeable
	// TODO: check to see if files in folder are moveable
	// TODO: throw meaningful errors
	_, err := os.ReadDir(folder)

	if err != nil {
		return false, err
	}

	return true, nil
}
