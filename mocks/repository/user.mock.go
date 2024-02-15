package repository_mock

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/ubaidillahhf/dating-service/app/domain"
	domain_mock "github.com/ubaidillahhf/dating-service/mocks/domain"
	"gorm.io/gorm"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) MockInsertSuccess() {
	m.Mock.On("Create", mock.Anything, mock.Anything).Return(domain_mock.MakeMockUser(), nil)
}

func (m *MockUserRepository) Create(ctx context.Context, newData domain.User) (domain.User, error) {
	args := m.Called(ctx, newData)

	var e error
	var r domain.User

	if n, ok := args.Get(0).(domain.User); ok {
		r = n
	}

	if n, ok := args.Get(1).(error); ok {
		e = n
	}

	return r, e
}

func (m *MockUserRepository) MockFindByIdentifierSuccess(user domain.User) {
	m.Mock.On("FindByIdentifier", mock.Anything, mock.Anything, mock.Anything).Return(user, nil)
}

func (m *MockUserRepository) MockFindByIdentifierNil(user domain.RegisterRequest) {
	m.Mock.On("FindByIdentifier", mock.Anything, mock.Anything, mock.Anything).Return(domain.User{}, nil)
}

func (m *MockUserRepository) FindByIdentifier(ctx context.Context, username, email string) (res domain.User, err error) {
	args := m.Called(ctx, username, email)

	var e error
	r := domain.User{}

	if n, ok := args.Get(0).(domain.User); ok {
		r = n
	}

	if n, ok := args.Get(1).(error); ok {
		e = n
	}

	return r, e
}

func (m *MockUserRepository) MockFindSuccess(user domain.User) {
	m.Mock.On("Find", mock.Anything, mock.Anything).Return(user, nil)
}

func (m *MockUserRepository) MockFindNil(user domain.RegisterRequest) {
	m.Mock.On("Find", mock.Anything, mock.Anything).Return(domain.User{}, nil)
}

func (m *MockUserRepository) Find(ctx context.Context, username string) (res domain.User, err error) {
	args := m.Called(ctx, username)

	var e error
	r := domain.User{}

	if n, ok := args.Get(0).(domain.User); ok {
		r = n
	}

	if n, ok := args.Get(1).(error); ok {
		e = n
	}

	return r, e
}

func (m *MockUserRepository) MockGetSuccess(user domain.User) {
	m.Mock.On("Get", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(user, nil)
}

func (m *MockUserRepository) Get(ctx context.Context, meta domain.Meta, myId string, excludeSeenDaily bool) (res []domain.User, total int64, err error) {
	args := m.Called(ctx, meta, myId)

	var e error
	r := []domain.User{}

	if n, ok := args.Get(0).([]domain.User); ok {
		r = n
	}

	if n, ok := args.Get(1).(int64); ok {
		total = n
	}

	if n, ok := args.Get(2).(error); ok {
		e = n
	}

	return r, total, e
}

func (m *MockUserRepository) MockSenderReceiverValidationSuccess(user domain.User) {
	m.Mock.On("SenderReceiverValidation", mock.Anything, mock.Anything, mock.Anything).Return(user, nil)
}

func (m *MockUserRepository) SenderReceiverValidation(ctx context.Context, senderId string, receiverId string) (res bool, err error) {
	args := m.Called(ctx, senderId, receiverId)

	var e error

	if n, ok := args.Get(0).(bool); ok {
		res = n
	}

	if n, ok := args.Get(1).(error); ok {
		e = n
	}

	return res, e
}

func (m *MockUserRepository) MockUpdateSuccess(user domain.User) {
	m.Mock.On("Update", mock.Anything, mock.Anything).Return(user, nil)
}

func (m *MockUserRepository) Update(ctx context.Context, newData domain.User) (res bool, err error) {
	args := m.Called(ctx, newData)

	var e error

	if n, ok := args.Get(0).(bool); ok {
		res = n
	}

	if n, ok := args.Get(1).(error); ok {
		e = n
	}

	return res, e
}

func (m *MockUserRepository) MockUpdateTxSuccess(user domain.User) {
	m.Mock.On("UpdateTx", mock.Anything, mock.Anything).Return(user, nil)
}

func (m *MockUserRepository) UpdateTx(ctx context.Context, tx *gorm.DB, newData domain.User) (res bool, err error) {
	args := m.Called(ctx, newData)

	var e error

	if n, ok := args.Get(0).(bool); ok {
		res = n
	}

	if n, ok := args.Get(1).(error); ok {
		e = n
	}

	return res, e
}
