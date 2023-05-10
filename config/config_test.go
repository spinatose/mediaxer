package config

import (
	"os"
	"testing"
)

const configFile string = "testConfig.json"

// Return new config object pointer with config defaults.
func TestNewConfig(t *testing.T) {
	// ARRANGE
	var config *Config
	
	// ACT
	config = NewConfig()

	// ASSERT
	if !config.MoveSourceFiles {
		t.Errorf("TestNewConfig: new configuration object returned with wrong value for MoveSourcefiles: %v", config.MoveSourceFiles)
	}

	if config.DestinationFolder != "" {
		t.Errorf("TestNewConfig: new configuration object returned with wrong value for DestinationFolder: %s", config.DestinationFolder)
	}
	
	if config.ExcludeSubFolders != false {
		t.Errorf("TestNewConfig: new configuration object returned with wrong value for ExcludeSubFolders: %v", config.ExcludeSubFolders)
	}
	
	if config.FileExtensionFilter != "*.txt" {
		t.Errorf("TestNewConfig: new configuration object returned with wrong value for FileExtensionFilter: %s", config.FileExtensionFilter)
	}

	if config.ResultFolderPattern != "YYYY_MMDD" {
		t.Errorf("TestNewConfig: new configuration object returned with wrong value for ResultFolderPattern: %s", config.ResultFolderPattern)
	}

	if config.SourceFolder != "./process" {
		t.Errorf("TestNewConfig: new configuration object returned with wrong value for SourceFolder: %s", config.SourceFolder)
	}
}

func TestCreateConfigFile(t *testing.T) {
	// ARRANGE
	var config *Config
	config = NewConfig()

	// ACT
	err := createConfigFileAndCleanup(t, config, configFile)

	// ASSERT
	if err != nil {
		t.Errorf("TestCreateConfigFile: failed to create config file [%s]- error: %s", configFile, err.Error())
	}
}

func TestLoadConfigFromJsonFile(t *testing.T) {
	// ARRANGE
	var config *Config
	config = NewConfig()
	config.SourceFolder = "spiny"

	if err := createConfigFileAndCleanup(t, config, configFile); err != nil {
		t.Errorf("TestLoadConfigFromJsonFile: unable to get new config for test- error: %s", err.Error())
	}

	// ACT
	newConfig, err := LoadConfigFromJsonFile(configFile)

	// ASSERT
	if err != nil {
		t.Errorf("TestLoadConfigFromJsonFile: failed to load config file [%s]- error: %s", configFile, err.Error())
	}

	if newConfig.SourceFolder != config.SourceFolder {
		t.Errorf("TestLoadConfigFromJsonFile: config file [%s] was not loaded properly- values don't correspond"+
			"config to config object originally created for test", configFile)
	}
}

func createConfigFileAndCleanup(t *testing.T, config *Config, configFile string) error {
	err := CreateConfigFile(config, configFile)
	t.Cleanup(func() {
		os.Remove(configFile)
	})

	return err
}
