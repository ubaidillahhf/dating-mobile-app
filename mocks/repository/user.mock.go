package repository_mock

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/ubaidillahhf/dating-service/app/domain"
	"github.com/ubaidillahhf/dating-service/app/infra/exception"
	domain_mock "github.com/ubaidillahhf/dating-service/mocks/domain"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) MockInsertSuccess() {
	m.Mock.On("Insert", mock.Anything, mock.Anything).Return(domain_mock.MakeMockUser(), nil)
}

func (m *MockUserRepository) Insert(ctx context.Context, newData domain.User) (domain.User, *exception.Error) {
	args := m.Called(ctx, newData)

	var e *exception.Error
	var r domain.User

	if n, ok := args.Get(0).(domain.User); ok {
		r = n
	}

	if n, ok := args.Get(1).(error); ok {
		e = &exception.Error{
			Code: exception.IntenalError,
			Err:  n,
		}
	}

	return r, e
}

func (m *MockUserRepository) MockFindByIdentifierSuccess(user domain.User) {
	m.Mock.On("FindByIdentifier", mock.Anything, mock.Anything, mock.Anything).Return(user, nil)
}

func (m *MockUserRepository) MockFindByIdentifierNil(user domain.User) {
	m.Mock.On("FindByIdentifier", mock.Anything, mock.Anything, mock.Anything).Return(domain.User{}, nil)
}

func (m *MockUserRepository) FindByIdentifier(ctx context.Context, username, email string) (res domain.User, err *exception.Error) {
	args := m.Called(ctx, username, email)

	var e error
	r := domain.User{}

	if n, ok := args.Get(0).(domain.User); ok {
		r = n
	}

	if n, ok := args.Get(1).(error); ok {
		e = n
	}

	return r, &exception.Error{
		Code: exception.IntenalError,
		Err:  e,
	}
}
