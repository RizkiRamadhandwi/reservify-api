package repo_mock

import (
	"booking-room-app/entity/dto"
	"time"

	"github.com/stretchr/testify/mock"
)

type ReportRepoMock struct {
	mock.Mock
}

func (r *ReportRepoMock) List(startDate, endDate time.Time) ([]dto.ReportDto, error) {
	args := r.Called(startDate, endDate)
	return args.Get(0).([]dto.ReportDto), args.Error(1)
}
