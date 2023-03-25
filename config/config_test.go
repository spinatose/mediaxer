package configuration

import (
	"fmt"
	"os"
	"testing"
)

const configFile string = "testConfig.json"

// Return new config object pointer with config defaults.
func TestNewConfig(t *testing.T) {
	config := NewConfig()

	if !config.MoveSourceFiles {
		t.Error(fmt.Sprintf("TestNewConfig: new configuration object returned with wrong value for MoveSourcefiles: %v", config.MoveSourceFiles))
	}
}

func TestCreateConfigFile(t *testing.T) {
	config := NewConfig()
	err := createConfigFileAndCleanup(t, config, configFile)

	if err != nil {
		t.Error(fmt.Sprintf("TestCreateConfigFile: failed to create config file [%s]- error: %s", configFile, err.Error()))
	}
}

func TestLoadConfigFromJsonFile(t *testing.T) {
	config := NewConfig()
	config.SourceFolder = "spiny"

	if err := createConfigFileAndCleanup(t, config, configFile); err != nil {
		t.Error(fmt.Sprintf("TestLoadConfigFromJsonFile: unable to get new config for test- error: %s", err.Error()))
	}

	newConfig, err := LoadConfigFromJsonFile(configFile)

	if err != nil {
		t.Error(fmt.Sprintf("TestLoadConfigFromJsonFile: failed to load config file [%s]- error: %s", configFile, err.Error()))
	}

	if newConfig.SourceFolder != config.SourceFolder {
		t.Error(fmt.Sprintf("TestLoadConfigFromJsonFile: config file [%s] was not loaded properly- values don't correspond"+
			"config to config object originally created for test", configFile))
	}
}

func createConfigFileAndCleanup(t *testing.T, config *Config, configFile string) error {
	err := CreateConfigFile(config, configFile)
	t.Cleanup(func() {
		os.Remove(configFile)
	})

	return err
}
