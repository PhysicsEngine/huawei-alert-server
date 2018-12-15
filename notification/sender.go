package notification

import (
	"go.uber.org/zap"
	"io"
	"net/http"
)

type Sender struct {
	url    string
	logger *zap.SugaredLogger
}

func (sender *Sender) send(body io.Reader) (*http.Response, error) {
	sender.logger.Infof("send notification to url=%s, body=%s", sender.url, body)
	return http.Post(sender.url, "application/json", body)
}

func createSender(logger *zap.SugaredLogger, url string) *Sender {
	return &Sender{url, logger}
}

func CreateLineSender(logger *zap.SugaredLogger) *Sender {
	return createSender(logger, "https://maker.ifttt.com/trigger/huawei_alert_line/with/key/c9GxSBX5gGyKITjQTGsuwH")
}

func CreateSlackSender(logger *zap.SugaredLogger) *Sender {
	return createSender(logger, "https://maker.ifttt.com/trigger/huawei_alert/with/key/c9GxSBX5gGyKITjQTGsuwH")
}
