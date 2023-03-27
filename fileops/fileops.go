package fileops

import (
	"os"
)

func ValidMachineFolder(folder string) (bool, error) {
	_, err := os.ReadDir(folder)

	if err != nil {
		return false, err
	}

	return true, nil
}
