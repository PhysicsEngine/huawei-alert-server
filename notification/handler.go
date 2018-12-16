package notification

import (
	"github.com/PhysicsEngine/huawei-alert-server/config"
	"go.uber.org/zap"
)

type Handler struct {
	senders map[string]*Sender
}

func CreateHandler(logger *zap.SugaredLogger, env *config.Env) *Handler {
	var senders = make(map[string]*Sender)
	senders["slack"] = CreateSender(logger, env.SlackUrl)
	senders["line"] = CreateSender(logger, env.LineUrl)
	senders["twitter"] = CreateSender(logger, env.TwitterUrl)
	return &Handler{senders}
}

func (handler *Handler) Send(name string) error {
	_, err := handler.senders[name].send(nil)
	return err
}

func (handler *Handler) Contains(name string) bool {
	_, contains := handler.senders[name]
	return contains
}
