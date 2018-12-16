package notification

import (
	"github.com/PhysicsEngine/huawei-alert-server/config"
	"go.uber.org/zap"
	"testing"
)

func TestCreateHandler(t *testing.T) {
	zapLogger, _ := zap.NewProduction()
	logger := zapLogger.Sugar()
	env, _ := config.ReadFromEnv()
	handler := CreateHandler(logger, env)
	if handler == nil {
		t.Fatalf("sender is null")
	}
}

func TestContains(t *testing.T) {
	zapLogger, _ := zap.NewProduction()
	logger := zapLogger.Sugar()
	env, _ := config.ReadFromEnv()
	handler := CreateHandler(logger, env)
	if !handler.Contains("slack") {
		t.Fatalf("slack should be contained")
	}
	if !handler.Contains("line") {
		t.Fatalf("line should be contained")
	}
	if !handler.Contains("twitter") {
		t.Fatalf("twitter should be contained")
	}
	if handler.Contains("foo") {
		t.Fatalf("foo should not be contained")
	}
}
