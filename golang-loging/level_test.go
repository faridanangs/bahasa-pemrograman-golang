package golangloging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

// urutan di jalankannya kode loger level
// PanicLevel,
// FatalLevel,
// ErrorLevel,
// WarnLevel,
// InfoLevel,
// DebugLevel,
// TraceLevel,

func TestLevel(t *testing.T) {
	logger := logrus.New()
	logger.Warn("Warn log")
	logger.Info("Info log")
	logger.Debug("Debug log")
	logger.Trace("Trace log")
	logger.Error("Error log")
	logger.Fatal("Fatal log") // ketika ini di jalankan ia akan meberhentikakn kodenya langsung
	// logger.Panic("Panic log") jika kita jalankan ini di atas maka kode di bawahnya akan berhenti
}

// trace loging level
func TestLogingLevel(t *testing.T) {
	// untuk mengatasi supaya trace dan debug bisa di jalankan kita beri tahu levelStatusnya
	// secara default dia menggunakan InfoLevel dan hanya menjalankan yang ada di atas level
	logger := logrus.New()
	// di sini semua kode di atas trace di jalankan
	logger.SetLevel(logrus.TraceLevel)
	logger.Warn("Warn log")
	logger.Info("Info log")
	logger.Debug("Debug log")
	logger.Trace("Trace log")
	logger.Error("Error log")
}
