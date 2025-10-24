package server

import (
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/notification"
	"github.com/relaunch-cot/service-notification/handler"
)

type Servers struct {
	Notification pb.NotificationServiceServer
}

func (s *Servers) Inject(handler *handler.Handlers) {
	s.Notification = NewNotificationServer(handler)
}
