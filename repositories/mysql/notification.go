package mysql

import (
	"context"
	"time"

	pbBaseModels "github.com/relaunch-cot/lib-relaunch-cot/proto/base_models"
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
	GetNotification(ctx *context.Context, notificationId string) (*pb.GetNotificationResponse, error)
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

func (r *mysqlResource) GetNotification(ctx *context.Context, notificationId string) (*pb.GetNotificationResponse, error) {
	var notification pbBaseModels.Notification

	baseQuery := `SELECT n.notificationId, n.senderId, n.receiverId, n.title, n.content, n.type, n.createdAt FROM notifications n WHERE n.notificationId = ?`

	row := mysql.DB.QueryRowContext(*ctx, baseQuery, notificationId)
	err := row.Scan(&notification.NotificationId, &notification.SenderId, &notification.ReceiverId, &notification.Title, &notification.Content, &notification.Type, &notification.CreatedAt)
	if err != nil {
		return nil, status.Error(codes.Internal, "error with database. Details: "+err.Error())
	}

	response := &pb.GetNotificationResponse{
		Notification: &notification,
	}

	return response, nil
}

func NewMysqlRepository(client *mysql.Client) IMysqlNotification {
	return &mysqlResource{
		client: client,
	}
}
