package golangloging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	// untuk menambahkan field di dalam {"level":"info","msg":"Info Mazzehh","time":"2023-09-22T20:52:36+08:00"} ini
	// kita bisa menggunakan WithField dan di tambah nama jenis lognya
	// logger.WithField("Key", value interface{}).Info("Info Mazzehh")
	logger.WithField("name", "farid anang s").Info("Info Mazzehh")

	// kita juga bisa menambah lebih dari 1 field
	logger.WithField("name", "farid anang s").WithField("kls", 12).WithField("absen", 9).Info("Info Mazzehh")
}

// cara yang lebih mudah jika fieldnya lebih dari satu
func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	// untuk membuat fields yang banyak kita bisa menggunakan kode ini
	// logger.WithFields(logrus.Fields{).Info("Field nih Bozz")
	logger.WithFields(logrus.Fields{
		"name":   "farid anang s",
		"kelas":  12,
		"absem":  9,
		"gender": "pria",
	}).Info("Field nih Bozz")
}
