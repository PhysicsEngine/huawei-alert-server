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

func CreateSender(logger *zap.SugaredLogger, url string) *Sender {
	logger.Infof("create ifttt handler. url:: %s", url)
	return &Sender{url, logger}
}
