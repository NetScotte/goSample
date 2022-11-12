package mylogrus

import (
	log "github.com/sirupsen/logrus"
)

func Basic() {
	logger := log.New()
	logger.Info("basic info log")
	logger.Infof("%v", logger.Out)
	logger.Infof("%v", logger.Formatter)
}

func JsonLog() {
	logger := log.New()
	logger.Formatter = new(log.JSONFormatter)
	logger.Info("json info log")
}
