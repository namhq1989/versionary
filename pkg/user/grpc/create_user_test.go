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

type createUserTestSuite struct {
	suite.Suite
	handler     grpc.CreateUserHandler
	mockCtrl    *gomock.Controller
	mockUserHub *mockuser.MockUserHub
}

func (s *createUserTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *createUserTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserHub = mockuser.NewMockUserHub(s.mockCtrl)

	s.handler = grpc.NewCreateUserHandler(s.mockUserHub)
}

func (s *createUserTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *createUserTestSuite) Test_1_Success() {
	// mock data
	s.mockUserHub.EXPECT().
		FindUserByEmail(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	s.mockUserHub.EXPECT().
		CreateUser(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.CreateUser(ctx, &userpb.CreateUserRequest{
		Name:  "Test user",
		Email: "test@gmail.com",
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *createUserTestSuite) Test_2_Fail_EmailExisted() {
	// mock data
	s.mockUserHub.EXPECT().
		FindUserByEmail(gomock.Any(), gomock.Any()).
		Return(&domain.User{
			ID:    database.NewStringID(),
			Name:  "Test user",
			Email: "test@gmail.com",
		}, nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.CreateUser(ctx, &userpb.CreateUserRequest{
		Name:  "Test user",
		Email: "test@gmail.com",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.EmailAlreadyExisted, err)
}

func (s *createUserTestSuite) Test_2_Fail_InvalidEmail() {
	// mock data
	s.mockUserHub.EXPECT().
		FindUserByEmail(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	// call
	ctx := appcontext.NewGRPC(context.Background())
	resp, err := s.handler.CreateUser(ctx, &userpb.CreateUserRequest{
		Name:  "Test user",
		Email: "invalid email",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.InvalidEmail, err)
}

//
// END OF CASES
//

func TestCreateUserTestSuite(t *testing.T) {
	suite.Run(t, new(createUserTestSuite))
}
