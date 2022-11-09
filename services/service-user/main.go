package main

import (
	"context"
	"log"
	"net"

	"clientserver/common/config"
	"clientserver/common/model"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var localStorage *model.UserList

func init() {
	localStorage = new(model.UserList)
	localStorage.List = make([]*model.User, 0)
}

type UsersServer struct {
	model.UnimplementedUsersServer
}

func (u UsersServer) Register(ctx context.Context, req *model.User) (*empty.Empty, error) {
	localStorage.List = append(localStorage.List, req)
	log.Println("Register user", req.String())

	return new(empty.Empty), nil
}

func (u UsersServer) Update(ctx context.Context, req *model.User) (*empty.Empty, error) {
	for i, user := range localStorage.List {
		if user.Id == req.Id {
			localStorage.List[i].Id = req.Id
			localStorage.List[i].Name = req.Name
			localStorage.List[i].Gender = req.Gender
			localStorage.List[i].Password = req.Password

			break
		}
	}
	log.Println("Update user", req.String())

	return new(empty.Empty), nil
}

func (u UsersServer) Delete(ctx context.Context, req *model.UserDelete) (*empty.Empty, error) {
	for i, user := range localStorage.List {
		if user.Id == req.Id {
			localStorage.List = RemoveIndex(localStorage.List, i)

			break
		}
	}
	log.Println("Delete user", req.String())

	return new(empty.Empty), nil
}

func RemoveIndex(listPerson []*model.User, index int) []*model.User {
	return append(listPerson[:index], listPerson[index+1:]...)
}

func (u UsersServer) List(context.Context, *empty.Empty) (*model.UserList, error) {
	return localStorage, nil
}

func main() {
	srv := grpc.NewServer()
	userSrv := UsersServer{}
	model.RegisterUsersServer(srv, userSrv)

	log.Println("Starting User Server at ", config.SERVICE_USER_PORT)

	listener, err := net.Listen("tcp", config.SERVICE_USER_PORT)
	if err != nil {
		log.Fatalf("could not listen. Err: %+v\n", err)
	}
	log.Fatalln(srv.Serve(listener))
}
