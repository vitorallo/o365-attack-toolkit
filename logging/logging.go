package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log logrus.Logger

func NewLogger(debug_mode string) *logrus.Logger {
	Log.Out = os.Stdout
	Log.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: false,
	})
	if debug_mode != "" {
		switch debug_mode {
		case "trace":
			Log.SetLevel(logrus.TraceLevel)
		case "debug":
			Log.SetLevel(logrus.DebugLevel)
		case "error":
			Log.SetLevel(logrus.ErrorLevel)
		default:
			Log.Fatal("You must specify one of the admitted logging levels when using -d flag")
		}
	} else {
		Log.SetLevel(logrus.ErrorLevel)
	}
	Log.Debug("Logging engine set to: ", Log.GetLevel())
	return &Log
}

func GetLogger() *logrus.Logger {
	return &Log
}
