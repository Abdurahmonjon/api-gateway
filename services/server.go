package services

import (
	"fmt"
	"github.com/Abdurahmonjon/api-gateway/config"
	pb "github.com/Abdurahmonjon/api-gateway/genproto/gitlab.com/Abdurahmonjon/studentproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
)

type IServiceManager interface {
	TaskService() pb.StudentServiceClient
}

type serviceManager struct {
	taskService pb.StudentServiceClient
}

func (s *serviceManager) TaskService() pb.StudentServiceClient {
	return s.taskService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connTask, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.StudentServiceHost, conf.StudentServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	log.Println("keldi mana")

	serviceManager := &serviceManager{
		taskService: pb.NewStudentServiceClient(connTask),
	}

	return serviceManager, nil
}
