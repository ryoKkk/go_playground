package play

import "github.com/sirupsen/logrus"

func LogInfo() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: false,
	})
	log.Infoln("this is a test")
}
