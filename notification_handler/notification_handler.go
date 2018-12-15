package notification_handler

import (
	"go.uber.org/zap"
	"net/http"
	"io"
)

type NotificationHandler struct {
	url string
	logger   *zap.SugaredLogger
}

func Send(handler *NotificationHandler, body io.Reader) (*http.Response, error){
	handler.logger.Infof("send notification to url=%s, body=%s", handler.url, body)
	return http.Post(handler.url, "application/json", body)
}

func createHandler(logger *zap.SugaredLogger, url string) *NotificationHandler {
	return &NotificationHandler{url, logger}
}

func CreateLineHandler(logger *zap.SugaredLogger) *NotificationHandler {
	return createHandler(logger, "https://maker.ifttt.com/trigger/huawei_alert_line/with/key/c9GxSBX5gGyKITjQTGsuwH")
}

func CreateSlackHandler(logger *zap.SugaredLogger) *NotificationHandler {
	return createHandler(logger, "https://maker.ifttt.com/trigger/huawei_alert/with/key/c9GxSBX5gGyKITjQTGsuwH")
}