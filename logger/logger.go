package logger

import "log"

type Logger struct {
	logger   *log.Logger
	logLevel LogLevel
}

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelWarning
	LevelError
)

func (l LogLevel) Includes(level LogLevel) bool {
	return l <= level
}

func (l Logger) Log(format string, v ...any) {
	if l.logger != nil {
		l.logger.Printf(format, v...)
	}
}

func (l Logger) Debug(format string, v ...any) {
	if l.logLevel.Includes(LevelDebug) {
		l.Log("DEBUG: "+format, v...)
	}
}

func (l Logger) Warning(format string, v ...any) {
	if l.logLevel.Includes(LevelWarning) {
		l.Log("WARNING: "+format, v...)
	}
}

func (l Logger) Error(format string, v ...any) {
	if l.logLevel.Includes(LevelError) {
		l.Log("Error: "+format, v...)
	}
}

func New(logger *log.Logger, logLevel LogLevel) Logger {
	return Logger{logger: logger, logLevel: logLevel}
}
