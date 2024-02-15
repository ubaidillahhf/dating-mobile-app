package usecases_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/ubaidillahhf/dating-service/app/domain"
	. "github.com/ubaidillahhf/dating-service/app/usecases"
	domain_mock "github.com/ubaidillahhf/dating-service/mocks/domain"
	repository_mock "github.com/ubaidillahhf/dating-service/mocks/repository"
)

type MockIUserUseCase struct {
	mock.Mock
}

type UserUseCaseSuite struct {
	suite.Suite
	usecase  IUserUsecase
	userRepo repository_mock.MockUserRepository
	newData  domain.RegisterRequest
}

func (s *UserUseCaseSuite) SetupTest() {
	s.userRepo = repository_mock.MockUserRepository{}
	s.usecase = NewUserUsecase(&s.userRepo)
	s.newData = domain_mock.MakeMockUserRegister()
}

func (s *UserUseCaseSuite) Test_UserRegister() {
	s.userRepo.MockFindByIdentifierNil(s.newData)
	s.userRepo.MockInsertSuccess()

	res, err := s.usecase.Register(context.Background(), s.newData)
	s.NotNil(res)
	s.Nil(err)
}

func Test_IUserUseCase(t *testing.T) {
	suite.Run(t, new(UserUseCaseSuite))
}
