package logging

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"spinatose.com/mediaxer/config"
)

type LogLevel int

const (
	Trace LogLevel = iota
	Debug
	Info
	Warn
	Error
	Fatal
	Panic
)

type logger struct {
	logEntry *log.Entry
	logLevel LogLevel
}

type Logger struct {
	loggers []*logger
}

func NewLogger(configs []config.LogOutput) *Logger {
	globalLogger := &Logger{
		loggers: nil,
	}

	for i, lconfig := range configs {
		logEntry := log.WithFields(log.Fields{
			"application": "mediaxer",
			"level":       lconfig.Options.Level,
			"loggerId":    i,
			"output":      lconfig.LogType,
		})

		// Temporarily set log output to terminal
		logEntry.Logger.Out = os.Stdout

		// Set log level of log entry to Trace so that it would write everything, but actual LogLevel
		// will be controlled by Global Logger.
		log.SetLevel(log.TraceLevel)

		logger := &logger{
			logEntry: logEntry,
			logLevel: parseLogLevel(lconfig.Options.Level),
		}

		globalLogger.loggers = append(globalLogger.loggers, logger)
	}

	return globalLogger
}

func (l *Logger) writeEntry(methodLogLevel LogLevel, args interface{}){
	for _, logger := range l.loggers {
		if logger.logLevel >= methodLogLevel {
			switch(methodLogLevel){
			case Trace:
				logger.logEntry.Trace(args)
			case Debug:
				logger.logEntry.Debug(args)
			case Info:
				logger.logEntry.Info(args)
			case Warn:
				logger.logEntry.Warn(args)
			case Error:
				logger.logEntry.Error(args)
			case Fatal:
				logger.logEntry.Fatal(args)
			case Panic:
				logger.logEntry.Panic(args)
			}
		}
	}
}

func (l *Logger) Info(args interface{}) {
	l.writeEntry(Info, args)
}

func (l *Logger) Debug(args interface{}) {
	l.writeEntry(Debug, args)
}

func parseLogLevel(level string) LogLevel {
	level = strings.ToLower(strings.Trim(level, " "))

	switch level {
	case "trace":
		return Trace
	case "debug":
		return Debug
	case "info":
		return Info
	case "warn":
		return Warn
	case "error":
		return Error
	case "fatal":
		return Fatal
	case "panic":
		return Panic
	}

	return Info
}
