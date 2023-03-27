package main

import (
	"fmt"
	"spinatose.com/mediaxer/cmd"
)

// TODO: eventually add command line args to filter certain files
// TODO: allow running from config file instead of command line- or allow override 
// TODO: provide flag type sets for filters like "all photos", "all videos", "all media", "all text files", etc...
// TODO: allow for argument for source and target directories (will have to check both for valid folders)
// TODO: possibly allow to upload to cloud
// TODO: possibly create file streaming server to accept incoming "files" stream over network and save to target
// TODO: create GUI interface

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
