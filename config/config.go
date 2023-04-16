package config

import (
	"encoding/json"
	"fmt"
	"os"
)

//TODO- have config create processed folder if doesn't exist
const workingFolder string = "./process"

type Config struct {
	DestinationFolder   string `json:"destinationFolder"`
	FileExtensionFilter string `json:"fileExtensionFilter"`
	Logger 				Logger `json:"logger"`
	MoveSourceFiles     bool   `json:"moveSourceFiles"`
	ResultFolderPattern string `json:"resultFolderPattern"`
	SourceFolder        string `json:"sourceFolder"`
}

type Logger struct {
	Level    string          `json:"level"`	
	Outputs []LogOutput `json:"outputs"` 
}

type LogOutputConfig struct {
	Colorize bool            `json:"colorize"`
	Format   string          `json:"format"`
	Config   LogOptionConfig `json:"config"`
}

type LogOptionConfig struct {
	Path     string            `json:"path"`
	FileName string            `json:"filename"`
}

type LogOutput struct {
	LogType string          `json:"logtype"`
	Options LogOutputConfig `json:"options"`
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
		DestinationFolder:   "",
		FileExtensionFilter: "*.txt",
		Logger: Logger{ 
			Level: "debug",
			Outputs: []LogOutput {
				{ 
					LogType: "console",
					Options: LogOutputConfig{
						Colorize: true,
					},
				},
				{
					LogType: "file",
					Options: LogOutputConfig{
						Colorize: false,
						Config: LogOptionConfig {
							FileName: "run.log",
							Path: "./process",
						},
						Format: "text",
					},
				},
			},
		},
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
	returnString += "\tLoggers:\n"
	returnString += fmt.Sprintf("\t\tLevel:\t\t%v\n", config.Logger.Level)

	for outputInc, logOut := range config.Logger.Outputs {
		returnString += fmt.Sprintf("\t\tOutput%v:\n", outputInc + 1)
		returnString += fmt.Sprintf("\t\t\tLogType:\t%s\n", logOut.LogType)
		returnString += "\t\t\tOptions:\n"
		returnString += fmt.Sprintf("\t\t\t\tColorize:\t%v\n", logOut.Options.Colorize)
		returnString += fmt.Sprintf("\t\t\t\tFormat:\t\t%v\n", logOut.Options.Format)
		returnString += "\t\t\t\tConfig:\n"
		returnString += fmt.Sprintf("\t\t\t\t\tFileName:\t%v\n", logOut.Options.Config.FileName)
		returnString += fmt.Sprintf("\t\t\t\t\tPath:\t\t%v\n", logOut.Options.Config.Path)
	}

	return returnString 
}
