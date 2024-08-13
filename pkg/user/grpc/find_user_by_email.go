package grpc

import (
	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/versionary-server/internal/genproto/userpb"
	"github.com/namhq1989/versionary-server/pkg/user/domain"
)

type FindUserByEmailHandler struct {
	userHub domain.UserHub
}

func NewFindUserByEmailHandler(userHub domain.UserHub) FindUserByEmailHandler {
	return FindUserByEmailHandler{
		userHub: userHub,
	}
}

func (h FindUserByEmailHandler) FindUserByEmail(ctx *appcontext.AppContext, req *userpb.FindUserByEmailRequest) (*userpb.FindUserByEmailResponse, error) {
	ctx.Logger().Info("[hub] new find user by email request", appcontext.Fields{"email": req.GetEmail()})

	ctx.Logger().Text("find user by email in db")
	user, err := h.userHub.FindUserByEmail(ctx, req.GetEmail())
	if err != nil {
		ctx.Logger().Error("failed to find user by email in db", err, appcontext.Fields{})
		return nil, err
	}
	if user == nil {
		ctx.Logger().ErrorText("user not found")
		return &userpb.FindUserByEmailResponse{User: nil}, nil
	}

	ctx.Logger().Text("done find user by email request")
	return &userpb.FindUserByEmailResponse{
		User: &userpb.User{
			Id:   user.ID,
			Name: user.Name,
		},
	}, nil
}
