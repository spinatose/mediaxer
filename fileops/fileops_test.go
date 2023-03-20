package fileops

import (
	"testing"
)

func TestValidMachineFolder(t *testing.T) {
	folder := "not a valid folder at all"

	// should return that folder is invalid- false
	if ValidMachineFolder(folder) {
		t.Error("ValidMachineFolder returned 'true' for an invalid folder")
	}
}
