package xorm_logrus

import (
	"github.com/sirupsen/logrus"
	"xorm.io/xorm/log"
)

// LogrusLogger is the logrus implement of ILogger
type LogrusLogger struct {
	logger  *logrus.Logger
	level   log.LogLevel
	showSQL bool
}

func (ll *LogrusLogger) Debug(v ...interface{}) {
	ll.logger.Debug(v...)
}

func (ll *LogrusLogger) Debugf(format string, v ...interface{}) {
	ll.logger.Debugf(format, v...)
}

func (ll *LogrusLogger) Error(v ...interface{}) {
	ll.logger.Error(v...)
}

func (ll *LogrusLogger) Errorf(format string, v ...interface{}) {
	ll.logger.Errorf(format, v...)
}

func (ll *LogrusLogger) Info(v ...interface{}) {
	ll.logger.Info(v...)
}

func (ll *LogrusLogger) Infof(format string, v ...interface{}) {
	ll.logger.Infof(format, v...)
}

func (ll *LogrusLogger) Warn(v ...interface{}) {
	ll.logger.Warn(v...)
}

func (ll *LogrusLogger) Warnf(format string, v ...interface{}) {
	ll.logger.Warnf(format, v...)
}

func (ll *LogrusLogger) Level() log.LogLevel {
	return ll.level
}

func (ll *LogrusLogger) SetLevel(l log.LogLevel) {
	ll.level = l
}

func (ll *LogrusLogger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		ll.showSQL = true
		return
	}
	ll.showSQL = show[0]
}

func (ll *LogrusLogger) IsShowSQL() bool {
	return ll.showSQL
}

func NewLogrusLogger() log.Logger {
	return NewLogrusLogger2(logrus.New())
}
func NewLogrusLogger2(logger *logrus.Logger) log.Logger {
	return NewLogrusLogger3(logger, log.DEFAULT_LOG_LEVEL)
}
func NewLogrusLogger3(logger *logrus.Logger, level log.LogLevel) log.Logger {
	return &LogrusLogger{logger: logger, level: level}
}
