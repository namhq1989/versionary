package user

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/versionary-server/internal/monolith"
	"github.com/namhq1989/versionary-server/pkg/user/application"
	"github.com/namhq1989/versionary-server/pkg/user/grpc"
	"github.com/namhq1989/versionary-server/pkg/user/infrastructure"
	"github.com/namhq1989/versionary-server/pkg/user/rest"
)

type Module struct{}

func (Module) Name() string {
	return "USER"
}

func (Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	var (
		userRepository = infrastructure.NewUserRepository(mono.Database())

		userHub = infrastructure.NewUserHub(mono.Database())

		// app
		app = application.New(userRepository)
		hub = grpc.New(userHub)
	)

	// rest server
	if err := rest.RegisterServer(ctx, app, mono.Rest(), mono.JWT()); err != nil {
		return err
	}

	// grpc server
	if err := grpc.RegisterServer(ctx, mono.RPC(), hub); err != nil {
		return err
	}

	return nil
}
