package golangloging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLoger(t *testing.T) {
	// untuk membuat logrus kita cukup mengetik logrus.New()
	logger := logrus.New()
	logger.Println("Hello Logrus")
	fmt.Println("Hello fmt")
}
