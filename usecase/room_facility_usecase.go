package usecase

import (
	"booking-room-app/entity"
	"booking-room-app/repository"
	"booking-room-app/shared/model"
	"fmt"
)

type RoomFacilityUsecase interface {
	FindAllRoomFacility(page, size int) ([]entity.RoomFacility, model.Paging, error)
	FindRoomFacilityById(id string) (entity.RoomFacility, error)
	AddRoomFacilityTransaction(payload entity.RoomFacility) (entity.RoomFacility, error)
	UpdateRoomFacilityTransaction(payload entity.RoomFacility) (entity.RoomFacility, error)
}

type roomFacilityUsecase struct {
	repo repository.RoomFacilityRepository
}

// find all room-facility
func (rf *roomFacilityUsecase) FindAllRoomFacility(page int, size int) ([]entity.RoomFacility, model.Paging, error) {
	if page == 0 && size == 0 {
		page = 1
		size = 5
	}
	return rf.repo.ListRoomFacility(page, size)
}

// find room-facility by id
func (rf *roomFacilityUsecase) FindRoomFacilityById(id string) (entity.RoomFacility, error) {
	return rf.repo.GetRoomFacilityById(id)
}

// add room-facility
func (rf *roomFacilityUsecase) AddRoomFacilityTransaction(payload entity.RoomFacility) (entity.RoomFacility, error) {
	// Check that the quantity entered does not exceed the quantity in facility
	quantity, err := rf.repo.GetQuantityFacilityByID(payload.FacilityId)
	if err != nil {
		return entity.RoomFacility{}, err
	}
	if payload.Quantity > quantity {
		return entity.RoomFacility{}, fmt.Errorf("oppps, quantity exceeds the facility quantity")
	}
	newFacilityQuantity := quantity - payload.Quantity

	// create room-facility transaction
	transactions, err := rf.repo.CreateRoomFacility(payload, newFacilityQuantity)
	if err != nil {
		return entity.RoomFacility{}, fmt.Errorf("oppps, failed to save room-facility transations :%v", err.Error())
	}
	return transactions, nil
}

// update room-facility
func (rf *roomFacilityUsecase) UpdateRoomFacilityTransaction(payload entity.RoomFacility) (entity.RoomFacility, error) {
	// get old record
	oldRoomFacility, err := rf.repo.GetRoomFacilityById(payload.ID)
	if err != nil {
		return entity.RoomFacility{}, fmt.Errorf("oppps, failed to get previous data :%v", err.Error())
	}

	// partial update checking
	if payload.RoomId == "" {
		payload.RoomId = oldRoomFacility.RoomId
	}
	if payload.FacilityId == "" {
		payload.FacilityId = oldRoomFacility.FacilityId
	}
	newFacilityQuantity := -1 // -1 is integer for none changed quantity
	if payload.Quantity == 0 {
		payload.Quantity = oldRoomFacility.Quantity
	} else {
		// check that the quantity entered does not exceed the quantity in facility
		facilityQuantity, err := rf.repo.GetQuantityFacilityByID(payload.FacilityId)
		if err != nil {
			return entity.RoomFacility{}, err
		}
		newFacilityQuantity = oldRoomFacility.Quantity - payload.Quantity + facilityQuantity // surplus or defisit are included in this one formula
		if newFacilityQuantity < 0 {
			return entity.RoomFacility{}, fmt.Errorf("oppps, quantity exceeds the facility quantity")
		}
	}
	if payload.Description == "" {
		payload.Description = oldRoomFacility.Description
	}

	roomFacility, err := rf.repo.UpdateRoomFacility(payload, newFacilityQuantity)
	if err != nil {
		return entity.RoomFacility{}, fmt.Errorf("oppps, failed to update data transations :%v", err.Error())
	}
	return roomFacility, nil
}

func NewRoomFacilityUsecase(repo repository.RoomFacilityRepository) RoomFacilityUsecase {
	return &roomFacilityUsecase{repo: repo}
}
