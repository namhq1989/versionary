package monolith

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/versionary-server/internal/authentication"
	"github.com/namhq1989/versionary-server/internal/caching"
	"github.com/namhq1989/versionary-server/internal/config"
	"github.com/namhq1989/versionary-server/internal/database"
	"github.com/namhq1989/versionary-server/internal/monitoring"
	"github.com/namhq1989/versionary-server/internal/queue"
	appjwt "github.com/namhq1989/versionary-server/internal/utils/jwt"
	"github.com/namhq1989/versionary-server/internal/utils/waiter"
	"google.golang.org/grpc"
)

type Monolith interface {
	Config() config.Server
	Database() *database.Database
	Caching() *caching.Caching
	Rest() *echo.Echo
	RPC() *grpc.Server
	Waiter() waiter.Waiter
	JWT() *appjwt.JWT
	Monitoring() *monitoring.Monitoring
	Queue() *queue.Queue
	Authentication() *authentication.Authentication
}

type Module interface {
	Name() string
	Startup(ctx *appcontext.AppContext, monolith Monolith) error
}
