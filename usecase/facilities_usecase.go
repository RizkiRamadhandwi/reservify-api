package usecase

import (
	"booking-room-app/entity"
	"booking-room-app/repository"
	"booking-room-app/shared/model"
	"fmt"
)

type FacilitiesUseCase interface {
	FindAllFacilities(page, size int) ([]entity.Facilities, model.Paging, error)
	RegisterNewFacilities(payload entity.Facilities) (entity.Facilities, error)
	FindFacilitiesById(id string) (entity.Facilities, error)
	EditFacilities(payload entity.Facilities) (entity.Facilities, error)
}

type facilitiesUseCase struct {
	repo repository.FasilitiesRepository
}

// FindAllFacilities implements FacilitiesUseCase.
func (f *facilitiesUseCase) FindAllFacilities(page int, size int) ([]entity.Facilities, model.Paging, error) {
	return f.repo.List(page, size)
}

// FindFacilitiesById implements FacilitiesUseCase.
func (f *facilitiesUseCase) FindFacilitiesById(id string) (entity.Facilities, error) {
	return f.repo.GetById(id)
}

// RegisterNewFacilities implements FacilitiesUseCase.
func (f *facilitiesUseCase) RegisterNewFacilities(payload entity.Facilities) (entity.Facilities, error) {
	if payload.Name == "" || payload.Quantity <= 0 {
		return entity.Facilities{}, fmt.Errorf("oppps, required fields")
	}

	facility, err := f.repo.Create(payload)
	if err != nil {
		return entity.Facilities{}, fmt.Errorf("opps failed save facility : %v", err)
	}

	return facility, nil
}

// EditFacilitiesById implements FacilitiesUseCase.
func (f *facilitiesUseCase) EditFacilities(payload entity.Facilities) (entity.Facilities, error) {
	if payload.Name == "" || payload.Quantity <= 0 {
		return entity.Facilities{}, fmt.Errorf("oppps, required fields")
	}

	facility, err := f.repo.UpdateById(payload)
	fmt.Println(facility)
	if err != nil {
		return entity.Facilities{}, fmt.Errorf("opps failed save facility : %v", err)
	}

	return facility, nil
}

func NewFacilitiesUseCase(repo repository.FasilitiesRepository) FacilitiesUseCase {
	return &facilitiesUseCase{repo: repo}
}
