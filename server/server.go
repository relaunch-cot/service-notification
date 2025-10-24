package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/notification"
	"github.com/relaunch-cot/service-notification/handler"
)

type notificationResource struct {
	handler *handler.Handlers
	pb.UnimplementedNotificationServiceServer
}

func (r *notificationResource) SendNotification(ctx context.Context, in *pb.SendNotificationRequest) (*empty.Empty, error) {
	err := r.handler.Notification.SendNotification(&ctx, in)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (r *notificationResource) GetNotification(ctx context.Context, in *pb.GetNotificationRequest) (*pb.GetNotificationResponse, error) {
	getNotificationResponse, err := r.handler.Notification.GetNotification(&ctx, in.NotificationId)
	if err != nil {
		return nil, err
	}

	return getNotificationResponse, nil
}

func (r *notificationResource) GetAllNotificationsFromUser(ctx context.Context, in *pb.GetAllNotificationsFromUserRequest) (*pb.GetAllNotificationsFromUserResponse, error) {
	getAllNotificationsFromUserResponse, err := r.handler.Notification.GetAllNotificationsFromUser(&ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	return getAllNotificationsFromUserResponse, nil
}

func NewNotificationServer(handler *handler.Handlers) pb.NotificationServiceServer {
	return &notificationResource{handler: handler}
}
