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
	GetAllNotificationsFromUser(ctx *context.Context, userId string) (*pb.GetAllNotificationsFromUserResponse, error)
	DeleteNotification(ctx *context.Context, notificationId string) error
	DeleteAllNotificationsFromUser(ctx *context.Context, userId string) error
}

func (r *mysqlResource) SendNotification(ctx *context.Context, notificationId string, in *pb.SendNotificationRequest) error {
	createdAt := time.Now()

	baseQuery := `INSERT INTO notifications (notificationId, senderId, receiverId, title, content, type, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := mysql.DB.ExecContext(*ctx, baseQuery, notificationId, in.SenderId, in.ReceiverId, in.Title, in.Content, in.Type, createdAt)
	if err != nil {
		return status.Error(codes.Internal, "error with database. Details: "+err.Error())
	}

	return nil
}

func (r *mysqlResource) GetNotification(ctx *context.Context, notificationId string) (*pb.GetNotificationResponse, error) {
	var notification pbBaseModels.Notification

	baseQuery := `
SELECT
	n.notificationId,
	n.senderId, 
	n.receiverId, 
	n.title,
	n.content,
	n.type,
    s.name AS senderName,
	n.createdAt
FROM notifications n 
	JOIN users s ON userId = n.senderId
WHERE n.notificationId = ?`

	row := mysql.DB.QueryRowContext(*ctx, baseQuery, notificationId)
	err := row.Scan(&notification.NotificationId, &notification.SenderId, &notification.ReceiverId, &notification.Title, &notification.Content, &notification.Type, &notification.SenderName, &notification.CreatedAt)
	if err != nil {
		return nil, status.Error(codes.Internal, "error with database. Details: "+err.Error())
	}

	response := &pb.GetNotificationResponse{
		Notification: &notification,
	}

	return response, nil
}

func (r *mysqlResource) GetAllNotificationsFromUser(ctx *context.Context, userId string) (*pb.GetAllNotificationsFromUserResponse, error) {
	var notifications []*pbBaseModels.Notification

	baseQuery := `
SELECT 
	n.notificationId, 
    n.senderId, 
    n.receiverId, 
    n.title, 
    n.content, 
    n.type, 
    s.name AS senderName,
    n.createdAt 
FROM notifications n 
	JOIN users s ON userId = n.senderId
WHERE n.receiverId = ? 
ORDER BY n.createdAt DESC`

	rows, err := mysql.DB.QueryContext(*ctx, baseQuery, userId)
	if err != nil {
		return nil, status.Error(codes.Internal, "error with database. Details: "+err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var notification pbBaseModels.Notification
		err := rows.Scan(&notification.NotificationId, &notification.SenderId, &notification.ReceiverId, &notification.Title, &notification.Content, &notification.Type, &notification.SenderName, &notification.CreatedAt)
		if err != nil {
			return nil, status.Error(codes.Internal, "error with database. Details: "+err.Error())
		}
		notifications = append(notifications, &notification)
	}

	response := &pb.GetAllNotificationsFromUserResponse{
		Notifications: notifications,
	}

	return response, nil
}

func (r *mysqlResource) DeleteNotification(ctx *context.Context, notificationId string) error {
	queryValidate := `SELECT COUNT(*) FROM notifications WHERE notificationId = ?`
	var count int
	err := mysql.DB.QueryRowContext(*ctx, queryValidate, notificationId).Scan(&count)
	if err != nil {
		return status.Error(codes.Internal, "error with database. Details: "+err.Error())
	}

	if count == 0 {
		return status.Error(codes.NotFound, "notification already deleted or does not exist")
	}

	baseQuery := `DELETE FROM notifications WHERE notificationId = ?`
	_, err = mysql.DB.ExecContext(*ctx, baseQuery, notificationId)
	if err != nil {
		return status.Error(codes.Internal, "error with database. Details: "+err.Error())
	}

	return nil
}

func (r *mysqlResource) DeleteAllNotificationsFromUser(ctx *context.Context, userId string) error {
	queryValidateUser := `SELECT * FROM users WHERE userId = ?`

	rows, err := mysql.DB.QueryContext(*ctx, queryValidateUser, userId)
	if err != nil {
		return status.Error(codes.Internal, "error with database. Details: "+err.Error())
	}

	if !rows.Next() {
		return status.Error(codes.NotFound, "user does not exist")
	}

	queryValidate := `SELECT COUNT(*) FROM notifications WHERE receiverId = ?`
	var count int
	err = mysql.DB.QueryRowContext(*ctx, queryValidate, userId).Scan(&count)
	if err != nil {
		return status.Error(codes.Internal, "error with database. Details: "+err.Error())
	}

	if count == 0 {
		return status.Error(codes.NotFound, "no notifications found for this user")
	}

	baseQuery := `DELETE FROM notifications WHERE receiverId = ?`
	_, err = mysql.DB.ExecContext(*ctx, baseQuery, userId)
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
