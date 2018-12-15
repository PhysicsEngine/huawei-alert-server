package notification

import "go.uber.org/zap"

type Handler struct {
	senders map[string]*Sender
}

func Create(logger *zap.SugaredLogger) *Handler {
	var senders = make(map[string]*Sender)
	senders["slack"] = CreateSlackSender(logger)
	senders["line"] = CreateLineSender(logger)
	return &Handler{senders}
}

func (handler *Handler) Send(name string) error {
	_, err := handler.senders[name].send(nil)
	return err
}
