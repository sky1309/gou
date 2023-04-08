package log

import (
	"log"
	"testing"
)

func TestLogFile(t *testing.T) {
	logger, err := New("debug", "./", log.LstdFlags)
	if err != nil {
		t.Error(err)
		return
	}
	logger.Debug("output debug")
	logger.Info("output info")
	logger.Error("output error")
	logger.Fatal("output fatal")
}
