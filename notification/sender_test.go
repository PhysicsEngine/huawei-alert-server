package notification

import (
	"go.uber.org/zap"
	"testing"
)

// setup logger
var zapLogger, _ = zap.NewProduction()
var logger = zapLogger.Sugar()

func TestCreateSlackSender(t *testing.T) {
	sender := CreateSlackSender(logger)
	if sender == nil {
		t.Fatalf("sender is null")
	}
}

func TestCreateLineSender(t *testing.T) {
	sender := CreateLineSender(logger)
	if sender == nil {
		t.Fatalf("sender is null")
	}
}

func TestSendSlack(t *testing.T) {
	sender := CreateSlackSender(logger)
	_, err := sender.send(nil)
	if err != nil {
		t.Fatalf("create huawai matcher faild %s", err)
	}
}

func TestSendLine(t *testing.T) {
	sender := CreateLineSender(logger)
	_, err := sender.send(nil)
	if err != nil {
		t.Fatalf("create huawai matcher faild %s", err)
	}
}
