package repo_mock

import (
	"booking-room-app/entity"
	"booking-room-app/shared/model"
	"time"

	"github.com/stretchr/testify/mock"
)

type TransactionsRepoMock struct {
	mock.Mock
}

// Create implements repository.TransactionsRepository.
func (t *TransactionsRepoMock) Create(payload entity.Transaction) (entity.Transaction, error) {
	args := t.Called(payload)
	return args.Get(0).(entity.Transaction), args.Error(1)
}

// GetTransactionByEmployeId implements repository.TransactionsRepository.

// GetTransactionById implements repository.TransactionsRepository.
func (t *TransactionsRepoMock) GetTransactionById(id string) (entity.Transaction, error) {
	args := t.Called(id)
	return args.Get(0).(entity.Transaction), args.Error(1)
}

// UpdatePemission implements repository.TransactionsRepository.
func (t *TransactionsRepoMock) UpdatePemission(payload entity.Transaction) (entity.Transaction, error) {
	args := t.Called(payload)
	return args.Get(0).(entity.Transaction), args.Error(1)
}

func (t *TransactionsRepoMock) List(page, size int, startDate, endDate time.Time) ([]entity.Transaction, model.Paging, error) {
	args := t.Called(page, size, startDate, endDate)
	return args.Get(0).([]entity.Transaction), args.Get(1).(model.Paging), args.Error(2)

}

func (t *TransactionsRepoMock) GetTransactionByEmployeId(EmployeeId string, page, size int) ([]entity.Transaction, model.Paging, error) {
	args := t.Called(EmployeeId, page, size)
	return args.Get(0).([]entity.Transaction), args.Get(1).(model.Paging), args.Error(2)
}

