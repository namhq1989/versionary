package query_test

import (
	"context"
	"testing"

	"github.com/namhq1989/go-utilities/appcontext"
	"github.com/namhq1989/versionary-server/internal/database"
	mockuser "github.com/namhq1989/versionary-server/internal/mock/user"
	apperrors "github.com/namhq1989/versionary-server/internal/utils/error"
	"github.com/namhq1989/versionary-server/pkg/user/application/query"
	"github.com/namhq1989/versionary-server/pkg/user/domain"
	"github.com/namhq1989/versionary-server/pkg/user/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type getMeTestSuite struct {
	suite.Suite
	handler            query.GetMeHandler
	mockCtrl           *gomock.Controller
	mockUserRepository *mockuser.MockUserRepository
}

func (s *getMeTestSuite) SetupSuite() {
	s.setupApplication()
}

func (s *getMeTestSuite) setupApplication() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mockUserRepository = mockuser.NewMockUserRepository(s.mockCtrl)

	s.handler = query.NewGetMeHandler(s.mockUserRepository)
}

func (s *getMeTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

//
// CASES
//

func (s *getMeTestSuite) Test_1_Success() {
	// mock data
	var id = database.NewStringID()

	s.mockUserRepository.EXPECT().
		FindUserByID(gomock.Any(), gomock.Any()).
		Return(&domain.User{
			ID:   id,
			Name: "Test user",
		}, nil)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetMe(ctx, id, dto.GetMeRequest{})

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), id, resp.User.ID)
}

func (s *getMeTestSuite) Test_2_Fail_InvalidID() {
	// mock data
	s.mockUserRepository.EXPECT().
		FindUserByID(gomock.Any(), gomock.Any()).
		Return(nil, apperrors.User.InvalidUserID)

	// call
	ctx := appcontext.NewRest(context.Background())
	resp, err := s.handler.GetMe(ctx, "invalid id", dto.GetMeRequest{})

	assert.NotNil(s.T(), err)
	assert.Nil(s.T(), resp)
	assert.Equal(s.T(), apperrors.User.InvalidUserID, err)
}

//
// END OF CASES
//

func TestGetMeTestSuite(t *testing.T) {
	suite.Run(t, new(getMeTestSuite))
}
