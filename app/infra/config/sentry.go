package config

import (
	"github.com/getsentry/sentry-go"
)

func SentryInit(configuration IConfig) {
	sentry.Init(sentry.ClientOptions{
		Dsn: configuration.Get("SENTRY_DSN"),
	})
}
