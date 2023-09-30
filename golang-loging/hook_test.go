package golangloging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type SampleHook struct{}

// kita mengecek jika level yang di kirim saya dengan level yang kite tentukan maka eksekusi func Fire
func (s *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s *SampleHook) Fire(entry *logrus.Entry) error {
	// kita ambil data yang ada di dalam entry kemudian ita tampilkan di di sini sesuai dengan yang kita mau
	fmt.Println("Hello hook", entry.Level, entry.Message, entry.Data)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	// kita masukan data sample hook ke dalam logrus hook jika levelhook
	// sama dengan yang ada di dalam SampleHook struct maka dia akan di jalankan secara otomatis
	logger.AddHook(&SampleHook{})
	logger.SetLevel(logrus.TraceLevel)

	logger.Info("Hello Info")
	logger.Error("Hello Error")
}
