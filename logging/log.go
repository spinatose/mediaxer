package logging

import (
	log "github.com/sirupsen/logrus"
	"spinatose.com/mediaxer/config"
)

type Logger struct {
	Level 	string 
	logger  *log.Entry
}

func NewLogger(config config.LogOutput) *Logger {
	logger := log.WithFields(log.Fields{
		"animal": "walrus",
		"number": 1,
		"size":   10,
	})

	return &Logger{
		Level: config.Options.Level,
		logger: logger,
	}
}

func (l *Logger) Info(args interface{}) {
	l.logger.Info(args)
}