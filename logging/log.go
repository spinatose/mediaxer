package logging

import (
	log "github.com/sirupsen/logrus"
	"spinatose.com/mediaxer/config"
)

type Logger struct {
	loggers  []*log.Entry
}

func NewLogger(configs []config.LogOutput) *Logger {
	globalLogger := &Logger{
		loggers: nil,
	}
	
	for _, lconfig := range configs {
		logger := log.WithFields(log.Fields{
			"application": "mediaxer",
			"level": lconfig.Options.Level,
		})
		
		globalLogger.loggers = append(globalLogger.loggers, logger)
	}

	return globalLogger
}

func (l *Logger) Info(args interface{}) {
	for _, logger := range l.loggers {
		logger.Info(args)
	}
}