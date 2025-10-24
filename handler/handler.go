package handler

import (
	"context"

	"github.com/google/uuid"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/notification"
	"github.com/relaunch-cot/service-notification/repositories"
)

type INotificationHandler interface {
	SendNotification(ctx *context.Context, in *pb.SendNotificationRequest) error
	GetNotification(ctx *context.Context, notificationId string) (*pb.GetNotificationResponse, error)
	GetAllNotificationsFromUser(ctx *context.Context, userId string) (*pb.GetAllNotificationsFromUserResponse, error)
}
type resource struct {
	repositories *repositories.Repositories
}

func (r *resource) SendNotification(ctx *context.Context, in *pb.SendNotificationRequest) error {
	notificationId := uuid.New().String()
	err := r.repositories.Mysql.SendNotification(ctx, notificationId, in)
	if err != nil {
		return err
	}

	return nil
}

func (r *resource) GetNotification(ctx *context.Context, notificationId string) (*pb.GetNotificationResponse, error) {
	notification, err := r.repositories.Mysql.GetNotification(ctx, notificationId)
	if err != nil {
		return nil, err
	}

	return notification, nil
}

func (r *resource) GetAllNotificationsFromUser(ctx *context.Context, userId string) (*pb.GetAllNotificationsFromUserResponse, error) {
	notifications, err := r.repositories.Mysql.GetAllNotificationsFromUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

func NewNotificationHandler(repositories *repositories.Repositories) INotificationHandler {
	return &resource{
		repositories: repositories,
	}
}
