package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.Trace("This is Trace") // engga ditampilin
	logger.Debug("This is Debug") // engga ditampilin
	logger.Info("This is Info")
	logger.Warn("This is Warn")
	logger.Error("This is Error")

	// trace dan debug ga ditampilin karena default level buat ditampilin adalah info ke atas
}

func TestLoggerLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel) // bakal nampilin log dari level trace

	logger.Trace("This is Trace")
	logger.Debug("This is Debug")
	logger.Info("This is Info")
	logger.Warn("This is Warn")
	logger.Error("This is Error")
}
