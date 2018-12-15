package notification

import (
	"go.uber.org/zap"
	"io"
	"net/http"
)

type NotificationSender struct {
	url    string
	logger *zap.SugaredLogger
}

func Send(sender *NotificationSender, body io.Reader) (*http.Response, error) {
	sender.logger.Infof("send notification to url=%s, body=%s", sender.url, body)
	return http.Post(sender.url, "application/json", body)
}

func createSender(logger *zap.SugaredLogger, url string) *NotificationSender {
	return &NotificationSender{url, logger}
}

func CreateLineSender(logger *zap.SugaredLogger) *NotificationSender {
	return createSender(logger, "https://maker.ifttt.com/trigger/huawei_alert_line/with/key/c9GxSBX5gGyKITjQTGsuwH")
}

func CreateSlackSender(logger *zap.SugaredLogger) *NotificationSender {
	return createSender(logger, "https://maker.ifttt.com/trigger/huawei_alert/with/key/c9GxSBX5gGyKITjQTGsuwH")
}
