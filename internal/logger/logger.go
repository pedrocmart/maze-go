package logger

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func NewLogger(output, format string) logrus.FieldLogger {

	if strings.ToUpper(output) != "STDOUT" {
		logFile, err := os.OpenFile(output, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			logrus.WithError(err).Fatalf("can't open log file")
		}
		logrus.SetOutput(logFile)
	} else {
		logrus.SetOutput(os.Stdout)
	}

	switch strings.ToLower(format) {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{
			PrettyPrint: false,
		})
	default:
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
	}

	return logrus.NewEntry(logrus.StandardLogger())
}
