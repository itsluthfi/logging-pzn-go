package main

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type SampleHook struct {
}

func (s *SampleHook) Levels() []logrus.Level { // level ketika hook akan dieksekusi
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s *SampleHook) Fire(entry *logrus.Entry) error { // perintah yang dijalankan ketika hook terjadi di level yg ditentukan sebelumnya
	fmt.Println("Sample Hook", entry.Level, entry.Message)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SampleHook{})
	logger.SetLevel(logrus.TraceLevel)

	logger.Info("Hello Info")
	logger.Warn("Hello Warn")
	logger.Error("Hello Error")
	logger.Debug("Hello Debug")
}
