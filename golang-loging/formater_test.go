package golangloging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestFormater(t *testing.T) {
	logger := logrus.New()
	// jika kita tidak menyetel nilai formatnya maka defaultnya adalah
	// TextFormatter
	// ini defaulnya nya
	// logger.SetFormatter(&logrus.TextFormatter{})

	// ini kita ubah menjadi json
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.TraceLevel)

	logger.Info("Info")
	logger.Error("Error")
	logger.Warn("Warn")
	logger.Debug("Debug")
	logger.Trace("Trace")
}
