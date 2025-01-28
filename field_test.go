package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) { // field sebagai informasi tambahan dari log
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "luthfi").
		WithField("name", "Luthfi Hanif").
		Info("Hello logger")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "luthfi",
		"name":     "Luthfi Hanif",
	}).
		Info("Hello logger")
}
