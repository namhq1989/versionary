package grpc

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/versionary-server/internal/genproto/userpb"
	"github.com/namhq1989/versionary-server/pkg/user/domain"
)

type FindUserByIDHandler struct {
	userHub domain.UserHub
}

func NewFindUserByIDHandler(userHub domain.UserHub) FindUserByIDHandler {
	return FindUserByIDHandler{
		userHub: userHub,
	}
}

func (h FindUserByIDHandler) FindUserByID(ctx *appcontext.AppContext, req *userpb.FindUserByIDRequest) (*userpb.FindUserByIDResponse, error) {
	ctx.Logger().Info("[hub] new find user by id request", appcontext.Fields{"userID": req.GetId()})

	ctx.Logger().Text("find user by id in db")
	user, err := h.userHub.FindUserByID(ctx, req.GetId())
	if err != nil {
		ctx.Logger().Error("failed to find user by id in db", err, appcontext.Fields{})
		return nil, err
	}
	if user == nil {
		ctx.Logger().ErrorText("user not found")
		return &userpb.FindUserByIDResponse{User: nil}, nil
	}

	ctx.Logger().Text("done find user by id request")
	return &userpb.FindUserByIDResponse{
		User: &userpb.User{
			Id:   user.ID,
			Name: user.Name,
		},
	}, nil
}
