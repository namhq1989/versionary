package monitoring

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/labstack/echo/v4"
)

type Monitoring struct {
	sentry *sentry.Client
}

func Init(_ *echo.Echo, dsn, machine, environment string) *Monitoring {
	// skip if the "machine" is not set
	if machine == "" {
		fmt.Printf("⚡️ [monitoring]: machine is not set \n")
		return nil
	}

	opts := sentry.ClientOptions{
		Dsn:           dsn,
		Environment:   fmt.Sprintf("%s-%s", environment, machine),
		EnableTracing: false,
	}

	if err := sentry.Init(opts); err != nil {
		panic(err)
	}

	client, err := sentry.NewClient(opts)
	if err != nil {
		panic(err)
	}

	// recover
	defer sentry.Recover()

	fmt.Printf("⚡️ [monitoring]: connected \n")

	return &Monitoring{sentry: client}
}
