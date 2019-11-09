package p2pb2b

import (
	"io"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func SetOutput(out io.Writer) {
	log.SetOutput(out)
}
