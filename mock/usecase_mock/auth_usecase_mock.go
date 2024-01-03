package usecase_mock

import (
	"booking-room-app/entity/dto"

	"github.com/stretchr/testify/mock"
)

type AuthUseCaseMock struct {
	mock.Mock
}

func (m *AuthUseCaseMock) Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	args := m.Called(payload)
	return args.Get(0).(dto.AuthResponseDto), args.Error(1)
}
