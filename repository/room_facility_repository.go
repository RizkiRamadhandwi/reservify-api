package repository

// create room facilities (ADMIN) -POST
// get all room facilities (ADMIN) -GET
// get by ID room facilities (ADMIN) -GET
// update room facilites (ADMIN) -GET

import (
	"booking-room-app/config"
	"booking-room-app/entity"
	"booking-room-app/shared/model"
	"database/sql"
	"log"
	"math"
	"net/http"
)

type RoomFacilityRepository interface {
	CreateRoomFacility(payload entity.RoomFacility, newQuantity int) (entity.RoomFacility, int, error)
	ListRoomFacility(page, size int) ([]entity.RoomFacility, model.Paging, error)
	GetRoomFacilityById(id string) (entity.RoomFacility, int, error)
	UpdateRoomFacility(payload entity.RoomFacility, newFacilityQuantity int) (entity.RoomFacility, int, error)
	GetQuantityFacilityByID(id string) (int, int, error)
}

type roomFacilityRepository struct {
	db *sql.DB
}

// get quantity from facilities by facility ID
func (t *roomFacilityRepository) GetQuantityFacilityByID(id string) (int, int, error) {
	var quantity int
	err := t.db.QueryRow(config.GetQuantityFacilityByID, id).Scan(&quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("roomFacilityRepository.QueryGetQuantityFacilityByID:", err.Error())
			return 0, http.StatusBadRequest, err
		}
		log.Println("roomFacilityRepository.QueryGetQuantityFacilityByID:", err.Error())
		return 0, http.StatusInternalServerError, err
	}
	return quantity, http.StatusOK, nil
}

// get all room facilities (ADMIN) -GET
func (t *roomFacilityRepository) ListRoomFacility(page, size int) ([]entity.RoomFacility, model.Paging, error) {
	var roomFacilities []entity.RoomFacility
	offset := (page - 1) * size

	rows, err := t.db.Query(config.SelectRoomFacilityList, size, offset)
	if err != nil {
		log.Println("roomFacilityRepository.Query:", err.Error())
		return nil, model.Paging{}, err
	}

	for rows.Next() {
		var roomFacility entity.RoomFacility
		err = rows.Scan(
			&roomFacility.ID,
			&roomFacility.RoomId,
			&roomFacility.FacilityId,
			&roomFacility.Quantity,
			&roomFacility.Description,
			&roomFacility.CreatedAt,
			&roomFacility.UpdatedAt)
		if err != nil {
			log.Println("roomFacilityRepository.Rows.Next():", err.Error())
			return nil, model.Paging{}, err
		}
		roomFacilities = append(roomFacilities, roomFacility)
	}

	totalRows := 0
	if err := t.db.QueryRow(config.GetCountRoomFacility).Scan(&totalRows); err != nil {
		log.Println("roomFacilityRepository.QueryRowGetCountRoomFacility:", err.Error())
		return nil, model.Paging{}, err
	}

	paging := model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}

	return roomFacilities, paging, nil

}

// get by ID room facilities (ADMIN) -GET
func (t *roomFacilityRepository) GetRoomFacilityById(id string) (entity.RoomFacility, int, error) {
	var roomFacility entity.RoomFacility
	err := t.db.QueryRow(config.SelectRoomFacilityByID, id).Scan(
		&roomFacility.ID,
		&roomFacility.RoomId,
		&roomFacility.FacilityId,
		&roomFacility.Quantity,
		&roomFacility.Description,
		&roomFacility.CreatedAt,
		&roomFacility.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("roomFacilityRepository.QueryRowSelectRoomFacilityByID:", err.Error())
			return entity.RoomFacility{}, http.StatusBadRequest, err
		}
		log.Println("roomFacilityRepository.QueryRowSelectRoomFacilityByID:", err.Error())
		return entity.RoomFacility{}, http.StatusInternalServerError, err
	}
	return roomFacility, http.StatusOK, nil
}

// create room facilities (ADMIN) -POST
func (t *roomFacilityRepository) CreateRoomFacility(payload entity.RoomFacility, newFacilityQuantity int) (entity.RoomFacility, int, error) {
	var roomFacilities entity.RoomFacility

	// begin transaction
	tx, err := t.db.Begin()
	if err != nil {
		log.Println("roomFacilityRepository.BeginTransaction:", err.Error())
		return entity.RoomFacility{}, http.StatusInternalServerError, err
	}

	// insert data
	err = tx.QueryRow(
		config.InsertTrxRoomFacility,
		payload.RoomId,
		payload.FacilityId,
		payload.Quantity,
		payload.Description).
		Scan(
			&payload.ID,
			&payload.CreatedAt,
			&payload.UpdatedAt)
	if err != nil {
		log.Println("roomFacilityRepository.QueryInsertData:", err.Error())
		return entity.RoomFacility{}, http.StatusInternalServerError, err
	}

	// reduce quantity in facility
	_, err = tx.Exec(config.UpdateQuantityFacilityByID, newFacilityQuantity, payload.FacilityId)
	if err != nil {
		log.Println("roomFacilityRepository.QueryReduceQuantity:", err.Error())
		return entity.RoomFacility{}, http.StatusInternalServerError, err
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println("roomFacilityRepository.TransactionCommit:", err.Error())
		return entity.RoomFacility{}, http.StatusInternalServerError, err
	}

	roomFacilities = payload
	return roomFacilities, http.StatusOK, err
}

// update room facilites (ADMIN) -GET
func (t *roomFacilityRepository) UpdateRoomFacility(payload entity.RoomFacility, newFacilityQuantity int) (entity.RoomFacility, int, error) {
	var roomFacility entity.RoomFacility

	// begin transaction
	tx, err := t.db.Begin()
	if err != nil {
		log.Println("roomFacilityRepository.BeginTransaction:", err.Error())
		return entity.RoomFacility{}, http.StatusInternalServerError, err
	}

	// update room-facility
	err = tx.QueryRow(
		config.UpdateRoomFacility,
		payload.RoomId,
		payload.FacilityId,
		payload.Quantity,
		payload.Description,
		payload.ID).Scan(&payload.CreatedAt, &payload.UpdatedAt)
	if err != nil {
		log.Println("roomFacilityRepository.UpdateRoomFacility:", err.Error())
		return entity.RoomFacility{}, http.StatusInternalServerError, err
	}

	// change facility quantity
	if newFacilityQuantity != -1 {
		_, err = tx.Exec(config.UpdateQuantityFacilityByID, newFacilityQuantity, payload.FacilityId)
		if err != nil {
			log.Println("roomFacilityRepository.QueryChangeQuantity:", err.Error())
			return entity.RoomFacility{}, http.StatusInternalServerError, err
		}
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println("roomFacilityRepository.CommitTransaction:", err.Error())
		return entity.RoomFacility{}, http.StatusInternalServerError, err
	}

	roomFacility = payload
	return roomFacility, http.StatusCreated, err
}

func NewRoomFacilityRepository(db *sql.DB) RoomFacilityRepository {
	return &roomFacilityRepository{db: db}
}

