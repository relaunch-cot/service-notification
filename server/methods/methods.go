package methods

import (
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/notification"
	"github.com/relaunch-cot/service-notification/resource"
	"google.golang.org/grpc"
)

func RegisterGrpcServices(s *grpc.Server) {
	pb.RegisterNotificationServiceServer(s, resource.Server.Notification)
}
