package notification

import "go.uber.org/zap"

var m = make(map[string]*NotificationSender)

func Init(logger *zap.SugaredLogger) {
	m["slack"] = CreateSlackSender(logger)
	m["line"] = CreateLineSender(logger)
}

func getSender(name string) *NotificationSender{
	return m[name]
}
