package logo

import "github.com/sirupsen/logrus"

var log logrus.Entry

func SetInstance(logrusEntry logrus.Entry) {
	log = logrusEntry
}
