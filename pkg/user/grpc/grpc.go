package grpc

import (
	"context"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/versionary-server/internal/genproto/userpb"
	"google.golang.org/grpc"
)

type server struct {
	hub App
	userpb.UnimplementedUserServiceServer
}

var _ userpb.UserServiceServer = (*server)(nil)

func RegisterServer(_ *appcontext.AppContext, registrar grpc.ServiceRegistrar, hub *Application) error {
	userpb.RegisterUserServiceServer(registrar, server{hub: hub})
	return nil
}

func (s server) CreateUser(bgCtx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	return s.hub.CreateUser(appcontext.NewGRPC(bgCtx), req)
}

func (s server) FindUserByID(bgCtx context.Context, req *userpb.FindUserByIDRequest) (*userpb.FindUserByIDResponse, error) {
	return s.hub.FindUserByID(appcontext.NewGRPC(bgCtx), req)
}

func (s server) FindUserByEmail(bgCtx context.Context, req *userpb.FindUserByEmailRequest) (*userpb.FindUserByEmailResponse, error) {
	return s.hub.FindUserByEmail(appcontext.NewGRPC(bgCtx), req)
}
