// Package logrus package provides an adapter to the
// go-kit log.Logger interface.
package logrus

import (
	"errors"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	*logrus.Logger
}

var errMissingValue = errors.New("(MISSING)")

// NewLogrusLogger takes a *logrus.Logger and returns
// a logger that satisfies the go-kit log.Logger interface.
func NewLogrusLogger(logger *logrus.Logger) log.Logger {
	return &logrusLogger{logger}
}

func (l logrusLogger) Log(keyvals ...interface{}) error {
	fields := logrus.Fields{}
	for i := 0; i < len(keyvals); i += 2 {
		if i+1 < len(keyvals) {
			fields[fmt.Sprint(keyvals[i])] = keyvals[i+1]
		} else {
			fields[fmt.Sprint(keyvals[i])] = errMissingValue
		}
	}
	l.WithFields(fields).Info()
	return nil
}
