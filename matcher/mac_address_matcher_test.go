package matcher

import (
	"testing"
	"go.uber.org/zap"
)

// setup logger
var zapLogger, _ = zap.NewProduction()
var logger = zapLogger.Sugar()

func TestCreateHuaweiMatcher(t *testing.T) {
	_, err := createHuaweiMatcher(logger)
	if err != nil {
		t.Fatalf("create huawei matcher faild %s", err)
	}
}

func TestHuaweiMatch(t *testing.T) {
	matcher, err := createHuaweiMatcher(logger)
	if err != nil {
		t.Fatalf("create huawai matcher faild %s", err)
	}

	var target = "10-C6-1F-foobar"
	if matcher.match(target) == false {
		t.Fatalf("target=%s should be matched", target)
	}

	target = "foober"
	if matcher.match(target) {
		t.Fatalf("target=%s should not be matched", target)
	}
}
