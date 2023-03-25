package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

const workingFolder string = "~/tmp"

type Config struct {
	DestinationFolder   string `json:"destinationFolder"`
	FileExtensionFilter string `json:"fileExtensionFilter"`
	MoveSourceFiles     bool   `json:"moveSourceFiles"`
	ResultFolderPattern string `json:"resultFolderPattern"`
	SourceFolder        string `json:"sourceFolder"`
}

func CreateConfigFile(config *Config, configFile string) error {
	content, err := json.Marshal(&config)
	if err != nil {
		return err
	}

	err = os.WriteFile(configFile, content, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func LoadConfigFromJsonFile(configFile string) (*Config, error) {
	var config *Config

	fileConfig, err := os.Open(configFile)
	defer fileConfig.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	jsonParser := json.NewDecoder(fileConfig)
	jsonParser.Decode(&config)
	return config, nil
}

func NewConfig() *Config {
	return &Config{
		DestinationFolder:   workingFolder,
		FileExtensionFilter: "*.txt",
		MoveSourceFiles:     true,
		ResultFolderPattern: "YYYY_MMDD",
		SourceFolder:        workingFolder,
	}
}

func (config *Config) ToString() string {
	returnString := "Configuration:\n"
	returnString += fmt.Sprintf("\tDestinationFolder:\t\t%s\n", config.DestinationFolder)
	returnString += fmt.Sprintf("\tFileExtensionFilter:\t\t%s\n", config.FileExtensionFilter)
	returnString += fmt.Sprintf("\tMoveSourceFiles:\t\t%v\n", config.MoveSourceFiles)
	returnString += fmt.Sprintf("\tResultFolderPattern:\t\t%s\n", config.ResultFolderPattern)
	returnString += fmt.Sprintf("\tSourceFolder:\t\t\t%s\n", config.SourceFolder)

	return returnString 
}
