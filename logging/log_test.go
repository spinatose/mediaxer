package logging

import (
	"fmt"
	"testing"
	
	configuration "spinatose.com/mediaxer/config"
)

func TestMergeFieldMaps(t *testing.T) {
	// ARRANGE 
	first := make(map[string]interface{}, 2)
	first["cat1"] = "fluffy"
	first["cat2"] = "gordo"
	
	second := make(map[string]interface{}, 2)
	second["perro1"] = "lobo"
	second["perro2"] = "brownie"

	// ACT
	third := mergeFieldMaps(first, second)

	// ASSERT
	if len(third) != 4 {
		t.Error(fmt.Sprintf("TestMergeFieldMaps: failed to concatenate second map to first- combined length : %v", len(third)))
	}
}

func TestNewLogger(t *testing.T) {
	// ARRANGE
	config := configuration.NewConfig()

	// ACT 
	globalLogger := NewLogger(config.Logger.Outputs)

	// ASSERT
	if len(globalLogger.loggers) != 2 {
		t.Error(fmt.Sprintf("TestNewLogger: failed to new up appropriate default global logger- wrong number of loggers : %d", len(globalLogger.loggers)))
	}
}

func TestParseLogLevel(t *testing.T) {
	// ARRANGE
	var logLvlTrace, logLvlDebug, logLvlInfo, logLvlWarn LogLevel
	var logLvlError, logLvlFatal, logLvlPanic LogLevel

	// ACT
	logLvlTrace = parseLogLevel("trace")
	logLvlDebug = parseLogLevel("Debug")
	logLvlInfo = parseLogLevel("Info")
	logLvlWarn = parseLogLevel("WArn")
	logLvlError = parseLogLevel("Error")
	logLvlFatal = parseLogLevel("Fatal")
	logLvlPanic = parseLogLevel("PANIC")

	// ASSERT
	if logLvlTrace != Trace {
		t.Error("TestParseLogLevel failed to parse 'Trace' log level")
	}

	if logLvlDebug != Debug {
		t.Error("TestParseLogLevel failed to parse 'Debug' log level")
	}

	if logLvlInfo != Info {
		t.Error("TestParseLogLevel failed to parse 'Info' log level")
	}

	if logLvlWarn != Warn {
		t.Error("TestParseLogLevel failed to parse 'Warn' log level")
	}

	if logLvlError != Error {
		t.Error("TestParseLogLevel failed to parse 'Error' log level")
	}

	if logLvlFatal != Fatal {
		t.Error("TestParseLogLevel failed to parse 'Fatal' log level")
	}

	if logLvlPanic != Panic {
		t.Error("TestParseLogLevel failed to parse 'Panic' log level")
	}
}
