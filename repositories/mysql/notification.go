package mysql

import "github.com/relaunch-cot/lib-relaunch-cot/repositories/mysql"

type mysqlResource struct {
	client *mysql.Client
}

type IMysqlNotification interface {
}

func NewMysqlRepository(client *mysql.Client) IMysqlNotification {
	return &mysqlResource{
		client: client,
	}
}
