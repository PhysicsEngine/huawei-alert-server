package notification

import (
	"go.uber.org/zap"
	"testing"
)


func TestCreateSlackSender(t *testing.T) {
	zapLogger, _ := zap.NewProduction()
	logger := zapLogger.Sugar()
	sender := CreateSlackSender(logger)
	if sender == nil {
		t.Fatalf("sender is null")
	}
}

func TestCreateLineSender(t *testing.T) {
	zapLogger, _ := zap.NewProduction()
	logger := zapLogger.Sugar()
	sender := CreateLineSender(logger)
	if sender == nil {
		t.Fatalf("sender is null")
	}
}

func TestSendSlack(t *testing.T) {
	zapLogger, _ := zap.NewProduction()
	logger := zapLogger.Sugar()
	sender := CreateSlackSender(logger)
	_, err := sender.send(nil)
	if err != nil {
		t.Fatalf("create huawai matcher faild %s", err)
	}
}

func TestSendLine(t *testing.T) {
	zapLogger, _ := zap.NewProduction()
	logger := zapLogger.Sugar()
	sender := CreateLineSender(logger)
	_, err := sender.send(nil)
	if err != nil {
		t.Fatalf("create huawai matcher faild %s", err)
	}
}
