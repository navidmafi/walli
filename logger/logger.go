package logger

import (
	"os"

	"github.com/charmbracelet/log"
)

var Logger *log.Logger

func Init() {
	Logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: false,
		Level:           log.DebugLevel,
		Formatter:       log.TextFormatter,
		ReportCaller:    true,
	})
}
