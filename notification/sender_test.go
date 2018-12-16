package notification

import (
	"github.com/PhysicsEngine/huawei-alert-server/config"
	"go.uber.org/zap"
	"testing"
)

func TestCreateSlackSender(t *testing.T) {
	zapLogger, _ := zap.NewProduction()
	logger := zapLogger.Sugar()
	env, _ := config.ReadFromEnv()
	sender := CreateSender(logger, env.SlackUrl)
	if sender == nil {
		t.Fatalf("sender is null")
	}
}

func TestCreateLineSender(t *testing.T) {
	zapLogger, _ := zap.NewProduction()
	logger := zapLogger.Sugar()
	env, _ := config.ReadFromEnv()
	sender := CreateSender(logger, env.LineUrl)
	if sender == nil {
		t.Fatalf("sender is null")
	}
}

func TestTwitterSend(t *testing.T) {
	zapLogger, _ := zap.NewProduction()
	logger := zapLogger.Sugar()
	env, _ := config.ReadFromEnv()
	sender := CreateSender(logger, env.TwitterUrl)
	_, err := sender.send(nil)
	if err != nil {
		t.Fatalf("create huawai matcher faild %s", err)
	}
}

func TestSendLine(t *testing.T) {
	zapLogger, _ := zap.NewProduction()
	logger := zapLogger.Sugar()
	env, _ := config.ReadFromEnv()
	sender := CreateSender(logger, env.LineUrl)
	_, err := sender.send(nil)
	if err != nil {
		t.Fatalf("create huawai matcher faild %s", err)
	}
}

func TestSendSlack(t *testing.T) {
	zapLogger, _ := zap.NewProduction()
	logger := zapLogger.Sugar()
	env, _ := config.ReadFromEnv()
	sender := CreateSender(logger, env.SlackUrl)
	_, err := sender.send(nil)
	if err != nil {
		t.Fatalf("create huawai matcher faild %s", err)
	}
}
