package handler

import "github.com/relaunch-cot/service-notification/repositories"

type Handlers struct {
	Notification INotificationHandler
}

func (h *Handlers) Inject(repositories *repositories.Repositories) {
	h.Notification = NewNotificationHandler(repositories)
}
