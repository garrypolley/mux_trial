package logging

import (
	"github.com/Sirupsen/logrus"
	"github.com/evalphobia/logrus_sentry"
)

var Logger = logrus.New()

func init() {
	sentryLogHandler, err := logrus_sentry.NewSentryHook("", []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	})

	if err == nil {
		Logger.Hooks.Add(sentryLogHandler)
	}
}
