package main

import (
	"clientserver/common/config"
	"clientserver/common/model"
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func serviceGarage() model.GaragesClient {
	port := config.SERVICE_GARAGE_PORT
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	return model.NewGaragesClient(conn)
}

func serviceUser() model.UsersClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	return model.NewUsersClient(conn)
}

func main() {
	userSvc := serviceUser()
	garageSvc := serviceGarage()
	ctx := context.Background()

	user1 := &model.User{
		Id:       "u001",
		Name:     "Hafiz arrahman",
		Password: "lohhehlohheh",
		Gender:   model.UserGender_MALE,
	}
	_, _ = userSvc.Register(ctx, user1)

	user2 := &model.User{
		Id:       "u002",
		Name:     "Kelpon MacGregor Wati",
		Password: "rahasia",
		Gender:   model.UserGender_FEMALE,
	}
	_, _ = userSvc.Register(ctx, user2)

	users, _ := userSvc.List(ctx, new(empty.Empty))
	log.Printf("List Users %+v\n", users.GetList())

	garage1 := &model.Garage{
		Id:   "g001",
		Name: "Kalimalang",
		Coordinate: &model.GarageCoordinate{
			Latitude:  -6.10,
			Longitude: 107.08,
		},
	}
	_, _ = garageSvc.Add(ctx, &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: garage1,
	})
	_, _ = garageSvc.Add(ctx, &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: garage1,
	})
	garages, _ := garageSvc.List(ctx, &model.GarageUserId{UserId: user1.Id})
	log.Printf("List garages %+v\n", garages.GetList())

}
