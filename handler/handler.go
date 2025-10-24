package handler

import "github.com/relaunch-cot/service-notification/repositories"

type INotificationHandler interface {
}
type resource struct {
	repositories *repositories.Repositories
}

func NewNotificationHandler(repositories *repositories.Repositories) INotificationHandler {
	return &resource{
		repositories: repositories,
	}
}
