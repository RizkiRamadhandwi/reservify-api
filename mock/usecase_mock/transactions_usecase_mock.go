package usecase_mock

import (
	"booking-room-app/entity"
	"booking-room-app/shared/model"
	"time"

	"github.com/stretchr/testify/mock"
)

type TransactionsUseCaseMock struct {
	mock.Mock
}

// FindAllTransactions implements usecase.TransactionsUsecase.
func (t *TransactionsUseCaseMock) FindAllTransactions(page , size int, startDate time.Time, endDate time.Time) ([]entity.Transaction, model.Paging, error) {
	args := t.Called(page, size, startDate, endDate)
	return args.Get(0).([]entity.Transaction), args.Get(1).(model.Paging), args.Error(2)
}

// FindTransactionsByEmployeeId implements usecase.TransactionsUsecase.
func (t *TransactionsUseCaseMock) FindTransactionsByEmployeeId(employeeId string, page, size int) ([]entity.Transaction, model.Paging, error) {
	args := t.Called(employeeId)
	return args.Get(0).([]entity.Transaction), args.Get(1).(model.Paging), args.Error(2)
}

// FindTransactionsById implements usecase.TransactionsUsecase.
func (t *TransactionsUseCaseMock) FindTransactionsById(id string) (entity.Transaction, error) {
	args := t.Called(id)
	return args.Get(0).(entity.Transaction), args.Error(1)
}

func (t *TransactionsUseCaseMock) RequestNewBookingRooms(payload entity.Transaction) (entity.Transaction, error) {
	args := t.Called(payload)
	return args.Get(0).(entity.Transaction), args.Error(1)
}

func (t *TransactionsUseCaseMock) AccStatusBooking(payload entity.Transaction) (entity.Transaction, error) {
	args := t.Called(payload)
	return args.Get(0).(entity.Transaction), args.Error(1)
}