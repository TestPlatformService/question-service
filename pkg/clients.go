package pkg

import (
	"log"
	"question/config"
	"question/genproto/group"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GroupServiceClient() (group.GroupServiceClient, error) {
	grpc, err := grpc.NewClient(config.LoadConfig().USER_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("GroupService bilan bog'lanilmadi")
		return nil, err
	}
	group := group.NewGroupServiceClient(grpc)
	return group, nil
}
