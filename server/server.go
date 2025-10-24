package server

import (
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/notification"
	"github.com/relaunch-cot/service-notification/handler"
)

type notificationResource struct {
	handler *handler.Handlers
	pb.UnimplementedNotificationServiceServer
}

func NewNotificationServer(handler *handler.Handlers) pb.NotificationServiceServer {
	return &notificationResource{handler: handler}
}
