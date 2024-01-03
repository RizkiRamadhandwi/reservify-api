package usecase_mock

import (
	"booking-room-app/entity/dto"

	"github.com/stretchr/testify/mock"
)

type ReportUseCaseMock struct {
	mock.Mock
}

func (r *ReportUseCaseMock) PrintAllReports(rangeParam string) ([]dto.ReportDto, error) {
	args := r.Called(rangeParam)
	return args.Get(0).([]dto.ReportDto), args.Error(1)
}
