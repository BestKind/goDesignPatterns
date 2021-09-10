package strategy

import "testing"

func TestNewLogManager(t *testing.T) {
	fileLogging := &FileLogging{}
	logManager := NewLogManager(fileLogging)

	logManager.Info()

	logManager.Error()

	DBlogging := &DBLogging{}
	logManager.Logging = DBlogging

	logManager.Info()
	logManager.Error()
}
