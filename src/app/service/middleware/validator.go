package middlewares

import (
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type CustomValidator struct {
	validator *validator.Validate
	msg       string
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewCustomValidator() *CustomValidator {
	cv := &CustomValidator{validator: validator.New(), msg: "rc"}

	return cv
}

// validator ~ membuat file log
func GenerateLoging(ip, level_error, nama_fungsi, querrys string, err *error) {
	var query string
	if querrys != "" {
		query = querrys
	}

	//! Logger
	loggers := logrus.New()
	fileLog, _ := os.OpenFile("log/LOG_"+time.Now().Local().Format("2006-Jan-02")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	loggers.SetOutput(fileLog)
	loggers.SetFormatter(&logrus.JSONFormatter{})

	switch strings.ToLower(level_error) {
	case "panic":
		loggers.WithFields(logrus.Fields{
			"URI":          nama_fungsi,
			"METHOD":       "",
			"STATUS":       "",
			"IP":           ip,
			"Host":         "",
			"RequestID":    "",
			"LogError":     "",
			"ResponseSize": "",
			"RESPONSE":     &err,
			"FROM":         "Function Logger",
			"QUERY":        &query,
		}).Panic("HIT FNCT <= " + nama_fungsi)

	case "fatal":
		loggers.WithFields(logrus.Fields{
			"URI":          nama_fungsi,
			"METHOD":       "",
			"STATUS":       "",
			"IP":           ip,
			"Host":         "",
			"RequestID":    "",
			"LogError":     "",
			"ResponseSize": "",
			"RESPONSE":     &err,
			"FROM":         "Function Logger",
			"QUERY":        &query,
		}).Fatal("HIT FNCT <= " + nama_fungsi)

	case "error":
		loggers.WithFields(logrus.Fields{
			"URI":          nama_fungsi,
			"METHOD":       "",
			"STATUS":       "",
			"IP":           ip,
			"Host":         "",
			"RequestID":    "",
			"LogError":     "",
			"ResponseSize": "",
			"RESPONSE":     &err,
			"FROM":         "Function Logger",
			"QUERY":        &query,
		}).Error("HIT FNCT <= " + nama_fungsi)

	case "warn", "warning":
		loggers.WithFields(logrus.Fields{
			"URI":          nama_fungsi,
			"METHOD":       "",
			"STATUS":       "",
			"IP":           ip,
			"Host":         "",
			"RequestID":    "",
			"LogError":     "",
			"ResponseSize": "",
			"RESPONSE":     &err,
			"FROM":         "Function Logger",
			"QUERY":        &query,
		}).Warn("HIT FNCT <= " + nama_fungsi)

	case "info":

		loggers.WithFields(logrus.Fields{
			"URI":          nama_fungsi,
			"METHOD":       "",
			"STATUS":       "",
			"IP":           ip,
			"Host":         "",
			"RequestID":    "",
			"LogError":     "",
			"ResponseSize": "",
			"RESPONSE":     "",
			"FROM":         "Function Logger",
			"QUERY":        &query,
		}).Info("HIT FNCT <= " + nama_fungsi)

	case "debug":
		loggers.WithFields(logrus.Fields{
			"URI":          nama_fungsi,
			"METHOD":       "",
			"STATUS":       "",
			"IP":           ip,
			"Host":         "",
			"RequestID":    "",
			"LogError":     "",
			"ResponseSize": "",
			"RESPONSE":     "",
			"FROM":         "Function Logger",
			"QUERY":        &query,
		}).Debug("HIT FNCT <= " + nama_fungsi)

	case "trace":
		loggers.WithFields(logrus.Fields{
			"URI":          nama_fungsi,
			"METHOD":       "",
			"STATUS":       "",
			"IP":           ip,
			"Host":         "",
			"RequestID":    "",
			"LogError":     "",
			"ResponseSize": "",
			"RESPONSE":     "",
			"FROM":         "Function Logger",
			"QUERY":        &query,
		}).Trace("HIT FNCT <= " + nama_fungsi)
	}
}
