package logging

import (
	"testing"
)

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
