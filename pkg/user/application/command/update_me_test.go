package command_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/versionary-server/internal/database"
	mockuser "github.com/namhq1989/versionary-server/internal/mock/user"
	apperrors "github.com/namhq1989/versionary-server/internal/utils/error"
	"github.com/namhq1989/versionary-server/pkg/user/application/command"
	"github.com/namhq1989/versionary-server/pkg/user/domain"
	"github.com/namhq1989/versionary-server/pkg/user/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type updateMeTestSuite struct {
	suite.Suite
	handler            command.UpdateMeHandler
	mockCtrl           *gomock.Controller
	mockUserRepository *mockuser.MockUserRepository
}

func (s *updateMeTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *updateMeTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserRepository = mockuser.NewMockUserRepository(s.mockCtrl)

	s.handler = command.NewUpdateMeHandler(s.mockUserRepository)
}

func (s *updateMeTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *updateMeTestSuite) Test_1_Success() {
	// mock data
	var id = database.NewStringID()

	s.mockUserRepository.EXPECT().
		FindUserByID(gomock.Any(), gomock.Any()).
		Return(&domain.User{
			ID:   id,
			Name: "Test user",
		}, nil)

	s.mockUserRepository.EXPECT().
		UpdateUser(gomock.Any(), gomock.Any()).
		Return(nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.UpdateMe(ctx, id, dto.UpdateMeRequest{
		Name: "John",
	})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
}

func (s *updateMeTestSuite) Test_2_Fail_InvalidID() {
	s.mockUserRepository.EXPECT().
		FindUserByID(gomock.Any(), gomock.Any()).
		Return(nil, apperrors.User.InvalidUserID)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.UpdateMe(ctx, "invalid id", dto.UpdateMeRequest{
		Name: "John",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.User.InvalidUserID, err)
}

func (s *updateMeTestSuite) Test_2_Fail_InvalidName() {
	s.mockUserRepository.EXPECT().
		FindUserByID(gomock.Any(), gomock.Any()).
		Return(&domain.User{
			ID:   database.NewStringID(),
			Name: "Test user",
		}, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.UpdateMe(ctx, "invalid id", dto.UpdateMeRequest{
		Name: "J",
	})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.Common.InvalidName, err)
}

//
// END OF CASES
//

func TestUpdateMeTestSuite(t *testing.T) {
	suite.Run(t, new(updateMeTestSuite))
}
