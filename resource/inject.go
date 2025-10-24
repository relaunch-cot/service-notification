package resource

import (
	"github.com/relaunch-cot/service-notification/handler"
	"github.com/relaunch-cot/service-notification/repositories"
	"github.com/relaunch-cot/service-notification/server"
)

var Repositories repositories.Repositories
var Handler handler.Handlers
var Server server.Servers

func Inject() {
	mysqlClient := OpenMysqlConn()

	Repositories.Inject(mysqlClient)
	Handler.Inject(&Repositories)
	Server.Inject(&Handler)
}
