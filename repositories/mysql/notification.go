package mysql

import (
	"context"
	"time"

	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/notification"
	"github.com/relaunch-cot/lib-relaunch-cot/repositories/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type mysqlResource struct {
	client *mysql.Client
}

type IMysqlNotification interface {
	SendNotification(ctx *context.Context, notificationId string, in *pb.SendNotificationRequest) error
}

func (r *mysqlResource) SendNotification(ctx *context.Context, notificationId string, in *pb.SendNotificationRequest) error {
	createdAt := time.Now()

	baseQuery := `INSERT INTO notifications (notificationId, senderId, receiverId, title, content, type,  createdAt) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := mysql.DB.ExecContext(*ctx, baseQuery, notificationId, in.SenderId, in.ReceiverId, in.Title, in.Content, in.Type, createdAt)
	if err != nil {
		return status.Error(codes.Internal, "error with database. Details: "+err.Error())
	}

	return nil
}

func NewMysqlRepository(client *mysql.Client) IMysqlNotification {
	return &mysqlResource{
		client: client,
	}
}
