package handler

import (
	"context"

	"github.com/google/uuid"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/notification"
	"github.com/relaunch-cot/service-notification/repositories"
)

type INotificationHandler interface {
	SendNotification(ctx *context.Context, in *pb.SendNotificationRequest) error
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

func NewNotificationHandler(repositories *repositories.Repositories) INotificationHandler {
	return &resource{
		repositories: repositories,
	}
}
