package golangloging

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOutput(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	file, _ := os.OpenFile("loging.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	logger.SetOutput(file)
	logger.Info("Info Masalahnya apa")
	logger.Panic("Panic apa")
	logger.Error("Errornya apa")
	logger.Trace("Trace apa")
	logger.Debug("Debug apa")
	logger.Fatal("Fatanya apa")
}
