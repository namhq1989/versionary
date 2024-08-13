package grpc_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/versionary-server/internal/database"
	"github.com/namhq1989/versionary-server/internal/genproto/userpb"
	mockuser "github.com/namhq1989/versionary-server/internal/mock/user"
	apperrors "github.com/namhq1989/versionary-server/internal/utils/error"
	"github.com/namhq1989/versionary-server/pkg/user/domain"
	"github.com/namhq1989/versionary-server/pkg/user/grpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type findUserByEmailTestSuite struct {
	suite.Suite
	handler     grpc.FindUserByEmailHandler
	mockCtrl    *gomock.Controller
	mockUserHub *mockuser.MockUserHub
}

func (s *findUserByEmailTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *findUserByEmailTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserHub = mockuser.NewMockUserHub(s.mockCtrl)

	s.handler = grpc.NewFindUserByEmailHandler(s.mockUserHub)
}

func (s *findUserByEmailTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *findUserByEmailTestSuite) Test_1_Success() {
	// mock data
	var (
		id    = database.NewStringID()
		email = "test@gmail.com"
	)
	s.mockUserHub.EXPECT().
		FindUserByEmail(gomock.Any(), gomock.Any()).
		Return(&domain.User{
			ID:    id,
			Name:  "Test user",
			Email: email,
		}, nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.FindUserByEmail(ctx, &userpb.FindUserByEmailRequest{
		Email: email,
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), id, resp.GetUser().GetId())
}

func (s *findUserByEmailTestSuite) Test_2_Fail_InvalidEmail() {
	// mock data
	email := "invalid email"

	s.mockUserHub.EXPECT().
		FindUserByEmail(gomock.Any(), gomock.Any()).
		Return(nil, apperrors.Common.InvalidEmail)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.FindUserByEmail(ctx, &userpb.FindUserByEmailRequest{
		Email: email,
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.InvalidEmail, err)
}

//
// END OF CASES
//

func TestFindUserByEmailTestSuite(t *testing.T) {
	suite.Run(t, new(findUserByEmailTestSuite))
}
