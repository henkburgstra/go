package servicehandler

import (
	"bitbucket.org/kardianos/service"
	"github.com/henkburgstra/go/logging"
)

type ServiceHandler struct {
	logging.LogHandler
	service service.Service
}

func NewServiceHandler(service service.Service) *ServiceHandler {
	serviceHandler := new(ServiceHandler)
	serviceHandler.service = service
	serviceHandler.LogHandler = *logging.NewLogHandler()
	return serviceHandler
}

func (h *ServiceHandler) Handle(logRecord *logging.LogRecord) {
	h.Emit(logRecord)
}

func (h *ServiceHandler) Emit(logRecord *logging.LogRecord) {
	msg := h.Format(logRecord)
	switch logRecord.GetLevel() {
	case logging.WARNING:
		h.service.Warning(msg)
	case logging.ERROR:
		h.service.Error(msg)
	default:
		h.service.Info(msg)
	}
}
