package logger

import (
	"bewell-backend-challenge/util/helpers/appjson"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	logLevelInfo    = "Info"
	logLevelWarning = "Warning"
	logLevelError   = "Error"
	debugLog        = "[Debug]: %s: %s: %s"
)

type Logger interface {
	Debug(msg string, fields ...interface{})
	Info(msg string, fields ...interface{})
	Warn(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
	SetTraceID(traceID string)
	GetTraceID() string
}

func SetFormatter(
	topic string,
	statusCode int,
	request interface{},
	_ interface{},
) logrus.Fields {
	return logrus.Fields{
		"app":        viper.GetString("APPLICATION_NAME"),
		"topic":      topic,
		"statusCode": statusCode,
		"request":    request,
	}
}

func New() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	return log
}

func Info(logKey string, message interface{}) {
	logDataMessage := LogFormatter{
		Level:   logLevelInfo,
		Time:    time.Now(),
		Message: message,
	}
	json, _ := appjson.Stringtify(logDataMessage)
	logData := fmt.Sprintf(
		debugLog,
		strings.ToUpper(logLevelInfo),
		logKey,
		json,
	)

	logrus.Info(logData)
}

func Error(logKey string, message interface{}) {
	logDataMessage := LogFormatter{
		Level:   logLevelError,
		Time:    time.Now(),
		Message: message,
	}
	json, _ := appjson.Stringtify(logDataMessage)
	logData := fmt.Sprintf(
		debugLog,
		strings.ToUpper(logLevelError),
		logKey,
		json,
	)

	logrus.Error(logData)
}

func Fatal(logKey string, message interface{}) {
	logDataMessage := LogFormatter{
		Level:   logLevelError,
		Time:    time.Now(),
		Message: message,
	}
	json, _ := appjson.Stringtify(logDataMessage)
	logData := fmt.Sprintf(
		debugLog,
		strings.ToUpper(logLevelError),
		logKey,
		json,
	)

	logrus.Fatal(logData)
}

func Warning(logKey string, message interface{}) {
	logDataMessage := LogFormatter{
		Level:   logLevelWarning,
		Time:    time.Now(),
		Message: message,
	}
	json, _ := appjson.Stringtify(logDataMessage)
	logData := fmt.Sprintf(
		debugLog,
		strings.ToUpper(logLevelWarning),
		logKey,
		json,
	)

	logrus.Warning(logData)
}

type LogFormatter struct {
	Level         string              `json:"level"`
	Time          time.Time           `json:"time"`
	IP            string              `json:"ip"`
	Method        string              `json:"method"`
	Path          string              `json:"path"`
	StatusCode    int                 `json:"status_code"`
	ExecutionTime time.Duration       `json:"execution_time"`
	Request       logFormatterRequest `json:"request"`
	Response      interface{}         `json:"response"`
	Message       interface{}         `json:"message"`
}

type logFormatterRequest struct {
	Body     interface{} `json:"body"`
	JSONBody interface{} `json:"json_body"`
}
