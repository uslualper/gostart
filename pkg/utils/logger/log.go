package logger

import (
	"fmt"
	logSchema "gostart/pkg/schema/logger"
)

type OwnerType string

const (
	System  OwnerType = "system"
	MongoDB OwnerType = "mongodb"
)

type LogType string

const (
	Error   LogType = "error"
	Warning LogType = "warning"
	Info    LogType = "info"
)

type LogHandlerType string

const (
	MongoLogHandler LogHandlerType = "mongo"
)

type LogHandler interface {
	Save(owner OwnerType, context logSchema.Context, typ LogType, level int)
}

var log *Logger

type Logger struct {
	handlers []LogHandler
	debug    bool
}

func Instance() *Logger {
	return log
}

func Context(message string, trace string, params ...map[string]string) logSchema.Context {
	return logSchema.Context{
		Message: message,
		Trace:   trace,
		Params:  getParams(params...),
	}
}

func getParams(params ...map[string]string) []logSchema.Param {
	var logParams []logSchema.Param

	for _, p := range params {
		for k, v := range p {
			logParams = append(logParams, logSchema.Param{
				Key:   k,
				Value: v,
			})
		}
	}

	return logParams
}

func InitLogger(debug bool, handlers ...LogHandlerType) *Logger {
	if log == nil {
		log = &Logger{
			handlers: getHandlers(handlers...),
			debug:    debug,
		}
	}

	return log
}

func getHandlers(handlers ...LogHandlerType) []LogHandler {
	var logHandlers []LogHandler

	for _, h := range handlers {
		switch h {
		case MongoLogHandler:
			logHandlers = append(logHandlers, &Mongo{})
		}
	}

	return logHandlers
}

func (l *Logger) Error(owner OwnerType, context logSchema.Context, level int) {
	l.log(owner, context, Error, level)
}

func (l *Logger) Warning(owner OwnerType, context logSchema.Context, level int) {
	l.log(owner, context, Warning, level)
}

func (l *Logger) Info(owner OwnerType, context logSchema.Context, level int) {
	l.log(owner, context, Info, level)
}

func (l *Logger) log(owner OwnerType, context logSchema.Context, typ LogType, level int) {
	for _, h := range l.handlers {
		h.Save(owner, context, typ, level)
	}

	if l.debug {
		fmt.Println("owner:", owner, "type:", typ, "level:", level, "context:", context)
	}
}
