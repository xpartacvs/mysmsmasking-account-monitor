package main

import (
	"mysmsmasking-account-monitor/packages/logger"
	"mysmsmasking-account-monitor/packages/worker"
)

func main() {
	logger.Log().Info().Msg("Application is starting...")

	if err := worker.Start(); err != nil {
		logger.Log().Err(err).Msg("Application has been terminated due to unexpected error.")
	}
}
