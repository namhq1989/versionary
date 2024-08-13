package grpc_test

import (
	"context"
	"testing"

	apperrors "github.com/namhq1989/versionary-server/internal/utils/error"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/versionary-server/internal/database"
	"github.com/namhq1989/versionary-server/internal/genproto/userpb"
	mockuser "github.com/namhq1989/versionary-server/internal/mock/user"
	"github.com/namhq1989/versionary-server/pkg/user/domain"
	"github.com/namhq1989/versionary-server/pkg/user/grpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type findUserByIDTestSuite struct {
	suite.Suite
	handler     grpc.FindUserByIDHandler
	mockCtrl    *gomock.Controller
	mockUserHub *mockuser.MockUserHub
}

func (s *findUserByIDTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *findUserByIDTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserHub = mockuser.NewMockUserHub(s.mockCtrl)

	s.handler = grpc.NewFindUserByIDHandler(s.mockUserHub)
}

func (s *findUserByIDTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *findUserByIDTestSuite) Test_1_Success() {
	// mock data
	var (
		id = database.NewStringID()
	)
	s.mockUserHub.EXPECT().
		FindUserByID(gomock.Any(), gomock.Any()).
		Return(&domain.User{
			ID:    id,
			Name:  "Test user",
			Email: "test@gmail.com",
		}, nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.FindUserByID(ctx, &userpb.FindUserByIDRequest{
		Id: id,
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), id, resp.GetUser().GetId())
}

func (s *findUserByIDTestSuite) Test_2_Fail_InvalidEmail() {
	// mock data
	id := "invalid email"

	s.mockUserHub.EXPECT().
		FindUserByID(gomock.Any(), gomock.Any()).
		Return(nil, apperrors.User.InvalidUserID)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.FindUserByID(ctx, &userpb.FindUserByIDRequest{
		Id: id,
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.User.InvalidUserID, err)
}

//
// END OF CASES
//

func TestFindUserByIDTestSuite(t *testing.T) {
	suite.Run(t, new(findUserByIDTestSuite))
}
