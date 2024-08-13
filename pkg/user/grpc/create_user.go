package grpc

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/versionary-server/internal/genproto/userpb"
	apperrors "github.com/namhq1989/versionary-server/internal/utils/error"
	"github.com/namhq1989/versionary-server/pkg/user/domain"
)

type CreateUserHandler struct {
	userHub domain.UserHub
}

func NewCreateUserHandler(userHub domain.UserHub) CreateUserHandler {
	return CreateUserHandler{
		userHub: userHub,
	}
}

func (h CreateUserHandler) CreateUser(ctx *appcontext.AppContext, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	ctx.Logger().Info("[hub] new create user request", appcontext.Fields{"name": req.GetName(), "email": req.GetEmail(), "timezone": req.GetTimezone()})

	ctx.Logger().Text("check email existence")
	user, err := h.userHub.FindUserByEmail(ctx, req.GetEmail())
	if err != nil {
		ctx.Logger().Error("failed to check email existence", err, appcontext.Fields{})
		return nil, err
	}
	if user != nil {
		ctx.Logger().Error("email already exists", nil, appcontext.Fields{})
		return nil, apperrors.Common.EmailAlreadyExisted
	}

	ctx.Logger().Text("create new user's model")
	user, err = domain.NewUser(req.GetName(), req.GetEmail(), req.GetTimezone())
	if err != nil {
		ctx.Logger().Error("failed to create new user's model", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("set data")
	_ = user.SetProvider(req.GetProviderSource(), req.GetProviderUid())

	ctx.Logger().Text("persist user in db")
	err = h.userHub.CreateUser(ctx, *user)
	if err != nil {
		ctx.Logger().Error("failed to persist user in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("enqueue tasks")
	if err = h.enqueueTasks(ctx, *user); err != nil {
		ctx.Logger().Error("failed to enqueue tasks", err, appcontext.Fields{})
	}

	ctx.Logger().Info("done create user request", appcontext.Fields{"userID": user.ID})
	return &userpb.CreateUserResponse{
		Id: user.ID,
	}, nil
}

func (CreateUserHandler) enqueueTasks(_ *appcontext.AppContext, _ domain.User) error {
	// TODO: DO SOMETHING WITH QUEUE
	return nil
}
