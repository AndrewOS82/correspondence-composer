package log

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	Trace = logrus.TraceLevel
	Debug = logrus.DebugLevel
	Info  = logrus.InfoLevel
	Warn  = logrus.WarnLevel
	Error = logrus.ErrorLevel
	Fatal = logrus.FatalLevel
	Panic = logrus.PanicLevel
)

var DefaultFormatter = &logrus.JSONFormatter{
	FieldMap: logrus.FieldMap{
		logrus.FieldKeyMsg:   "message",
		logrus.FieldKeyLevel: "status",
	},
}

type (
	Config struct {
		ServiceName  string
		CustomLogger *logrus.Logger
		LogLevel     *logrus.Level
		Formatter    logrus.Formatter
	}

	Logger interface {
		logrus.FieldLogger
		LogWithFields(logrus.Level, string, ...Fields)
		TraceWithFields(string, ...Fields)
		DebugWithFields(string, ...Fields)
		InfoWithFields(string, ...Fields)
		WarnWithFields(string, ...Fields)
		ErrorWithFields(error, ...Fields)
		ErrorWithDDFields(error, ...Fields)
		FatalErrorWithDDFields(error, ...Fields)
	}

	Fields map[string]interface{}

	logrusLogger struct {
		*logrus.Entry
	}
)

// New initializes a new instance of a logger.
func New(c Config) Logger {
	var baseLogger *logrus.Logger

	// default if baseLogger is absent
	if baseLogger = c.CustomLogger; baseLogger == nil {
		baseLogger = logrus.New()
	}

	if c.Formatter == nil {
		baseLogger.SetFormatter(DefaultFormatter)
	} else {
		baseLogger.SetFormatter(c.Formatter)
	}

	// allow for custom log level setting. by default, debug and trace level logs are disabled
	if c.LogLevel != nil {
		baseLogger.SetLevel(*c.LogLevel)
	}

	// adds base fields to all log entries
	baseEntry := baseLogger.WithFields(logrus.Fields{
		"service": c.ServiceName,
		"source":  "go",
	})

	return &logrusLogger{baseEntry}
}

func (l *logrusLogger) LogWithFields(level logrus.Level, msg string, fields ...Fields) {
	lw := l.withFields(fields)
	switch level {
	case Trace:
		lw.Traceln(msg)
	case Debug:
		lw.Debugln(msg)
	case Info:
		lw.Infoln(msg)
	case Warn:
		lw.Warnln(msg)
	case Error:
		lw.Errorln(msg)
	case Fatal:
		lw.Fatalln(msg)
	case Panic:
		lw.Panicln(msg)
	default:
		lw.Warnf("unknown log level: %s", level)
		lw.Warningln(msg)
	}
}

func (l *logrusLogger) TraceWithFields(msg string, fields ...Fields) {
	l.withFields(fields).Traceln(msg)
}

func (l *logrusLogger) DebugWithFields(msg string, fields ...Fields) {
	l.withFields(fields).Debugln(msg)
}

func (l *logrusLogger) InfoWithFields(msg string, fields ...Fields) {
	l.withFields(fields).Infoln(msg)
}

func (l *logrusLogger) WarnWithFields(msg string, fields ...Fields) {
	l.withFields(fields).Warnln(msg)
}

func (l *logrusLogger) ErrorWithFields(err error, fields ...Fields) {
	l.withFields(fields).Errorln(err)
}

// ErrorWithDDFields generates Datadog specific fields from an error
func (l *logrusLogger) ErrorWithDDFields(err error, fields ...Fields) {
	l.withDDErrorFields(err, fields).Errorln(err)
}

func (l *logrusLogger) FatalErrorWithDDFields(err error, fields ...Fields) {
	l.withDDErrorFields(err, fields).Fatalln(err)
}

func (l *logrusLogger) withFields(fields []Fields) *logrus.Entry {
	logFields := make(logrus.Fields)
	for _, f := range fields {
		for k, v := range f {
			// check for datadog reserved attributes, c.f.
			// https://docs.datadoghq.com/logs/log_configuration/attributes_naming_convention/#reserved-attributes
			switch k {
			case "host", "source", "status", "service", "trace_id", "message":
				fmt.Printf("log message field '%s' is a reserved attribute", k)
			}
			logFields[k] = v
		}
	}
	return l.WithFields(logFields)
}

// generates Datadog specific fields from error
// The structure of these fields are designed for datadog's parser
// https://docs.datadoghq.com/logs/log_collection/
func (l *logrusLogger) withDDErrorFields(err error, fields []Fields) *logrus.Entry {
	msg, cause, stack := parseError(err)

	errorFields := make(logrus.Fields)

	for _, f := range fields {
		for k, v := range f {
			errorFields[k] = v
		}
	}

	errorFields[logrus.ErrorKey] = map[string]string{
		"message": msg,
		"stack":   stack,
	}
	errorFields["cause"] = cause

	return l.WithFields(errorFields)
}

// parses errors pkg error to return a msg, cause, and stack trace
// see errors.WithStack.Format for formatting verbs
func parseError(err error) (string, string, string) {
	cause := fmt.Sprint(errors.Cause(err))
	stack := fmt.Sprintf("%+v", err)
	unTrimmedMessage := fmt.Sprintf("%s", err)
	return strings.TrimSuffix(unTrimmedMessage, fmt.Sprintf(": %s", cause)), cause, stack
}
