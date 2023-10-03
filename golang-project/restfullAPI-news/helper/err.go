package helper

import "github.com/sirupsen/logrus"

func IfLogingErr(err error, text string) {
	log := logrus.New()
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Info(text)
	}
}
