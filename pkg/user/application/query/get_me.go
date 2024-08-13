package query

import (
	"github.com/namhq1989/go-utilities/appcontext"
	apperrors "github.com/namhq1989/versionary-server/internal/utils/error"
	"github.com/namhq1989/versionary-server/pkg/user/domain"
	"github.com/namhq1989/versionary-server/pkg/user/dto"
)

type GetMeHandler struct {
	userRepository domain.UserRepository
}

func NewGetMeHandler(userRepository domain.UserRepository) GetMeHandler {
	return GetMeHandler{
		userRepository: userRepository,
	}
}

func (h GetMeHandler) GetMe(ctx *appcontext.AppContext, performerID string, _ dto.GetMeRequest) (*dto.GetMeResponse, error) {
	ctx.Logger().Info("[query] new get me request", appcontext.Fields{"performerID": performerID})

	ctx.Logger().Text("find user by id in db")
	user, err := h.userRepository.FindUserByID(ctx, performerID)
	if err != nil {
		ctx.Logger().Error("failed to find user by id in db", err, appcontext.Fields{})
		return nil, err
	}
	if user == nil {
		ctx.Logger().ErrorText("user not found")
		return nil, apperrors.User.UserNotFound
	}

	ctx.Logger().Text("done get me request")
	return &dto.GetMeResponse{
		User: dto.User{}.FromDomain(*user),
	}, nil
}
