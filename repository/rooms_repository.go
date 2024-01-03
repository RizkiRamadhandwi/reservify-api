package repository

import (
	"booking-room-app/config"
	"booking-room-app/entity"
	"booking-room-app/shared/model"
	"database/sql"
	"log"
	"math"
	"time"
)

type RoomRepository interface {
	Create(payload entity.Room) (entity.Room, error)
	Get(id string) (entity.Room, error)
	List(page, size int) ([]entity.Room, model.Paging, error)
	ListStatus(status string, page, size int) ([]entity.Room, model.Paging, error)
	Update(payload entity.Room) (entity.Room, error)
	UpdateStatus(payload entity.Room) (entity.Room, error)
}

type roomRepository struct {
	db *sql.DB
}

// Create implements RoomRepository.
func (r *roomRepository) Create(payload entity.Room) (entity.Room, error) {
	var room entity.Room
	err := r.db.QueryRow(config.InsertRoom, payload.Name, payload.RoomType, payload.Capacity, payload.Status).Scan(&room.ID, &room.CreatedAt, &room.UpdatedAt)
	if err != nil {
		log.Println("roomRepository.CreateQueryRow", err.Error())
		return entity.Room{}, err
	}

	room.Name = payload.Name
	room.RoomType = payload.RoomType
	room.Capacity = payload.Capacity
	room.Status = payload.Status

	return room, nil
}

// Get implements RoomRepository.
func (r *roomRepository) Get(id string) (entity.Room, error) {
	var room entity.Room
	err := r.db.QueryRow(config.SelectRoomByID, id).Scan(&room.ID, &room.Name, &room.RoomType, &room.Capacity, &room.Status, &room.CreatedAt, &room.UpdatedAt)
	if err != nil {
		log.Println("roomRepository.GetQueryRow", err.Error())
		return entity.Room{}, err
	}

	return room, nil
}

// List implements RoomRepository.
func (r *roomRepository) List(page, size int) ([]entity.Room, model.Paging, error) {
	var rooms []entity.Room
	offset := (page - 1) * size

	rows, err := r.db.Query(config.SelectRoomList, size, offset)
	if err != nil {
		log.Println("roomRepository.ListQuery", err.Error())
		return []entity.Room{}, model.Paging{}, err
	}

	for rows.Next() {
		var room entity.Room
		err := rows.Scan(&room.ID, &room.Name, &room.RoomType, &room.Capacity, &room.Status, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			log.Println("roomRepository.ListScan", err.Error())
			return []entity.Room{}, model.Paging{}, err
		}

		rooms = append(rooms, room)
	}

	totalRows := 0
	if err := r.db.QueryRow(config.SelectCountRoom).Scan(&totalRows); err != nil {
		return nil, model.Paging{}, err
	}

	paging := model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}

	return rooms, paging, nil
}

// List with Status implements RoomRepository.
func (r *roomRepository) ListStatus(status string, page, size int) ([]entity.Room, model.Paging, error) {
	var rooms []entity.Room
	offset := (page - 1) * size

	rows, err := r.db.Query(config.SelectRoomListStatus, status, size, offset)
	if err != nil {
		log.Println("roomRepository.ListQuery", err.Error())
		return []entity.Room{}, model.Paging{}, err
	}

	for rows.Next() {
		var room entity.Room
		err := rows.Scan(&room.ID, &room.Name, &room.RoomType, &room.Capacity, &room.Status, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			log.Println("roomRepository.ListScan", err.Error())
			return []entity.Room{}, model.Paging{}, err
		}

		rooms = append(rooms, room)
	}

	totalRows := 0
	if err := r.db.QueryRow(config.SelectCountRoomStatus, status).Scan(&totalRows); err != nil {
		return nil, model.Paging{}, err
	}

	paging := model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}

	return rooms, paging, nil
}

// Update implements RoomRepository.
func (r *roomRepository) Update(payload entity.Room) (entity.Room, error) {
	var room entity.Room
	room.ID = payload.ID
	payload.UpdatedAt = time.Now()

	err := r.db.QueryRow(config.UpdateRoomByID, room.ID, payload.Name, payload.RoomType, payload.Capacity, payload.Status).Scan(&room.CreatedAt, &room.UpdatedAt)
	if err != nil {
		log.Println("roomRepository.UpdateQueryRow", err.Error())
		return entity.Room{}, err
	}

	room.Name = payload.Name
	room.RoomType = payload.RoomType
	room.Capacity = payload.Capacity
	room.Status = payload.Status

	return room, nil
}

// UpdateStatus implements RoomRepository.
func (r *roomRepository) UpdateStatus(payload entity.Room) (entity.Room, error) {
	var room entity.Room
	room.ID = payload.ID
	payload.UpdatedAt = time.Now()

	err := r.db.QueryRow(config.UpdateRoomStatus, room.ID, payload.Status).Scan(&room.Name, &room.RoomType, &room.Capacity, &room.CreatedAt, &room.UpdatedAt)
	if err != nil {
		log.Println("roomRepository.UpdateStatusQueryRow", err.Error())
		return entity.Room{}, err
	}

	room.Status = payload.Status

	return room, nil
}

// create room (ADMIN) -GET
// get all rooms (ALL ROLE) -GET
// get by room by ID (ALL ROLE) -GET
// update room status (GA & ADMIN) -PUT
// update room (ADMIN) -PUT

func NewRoomRepository(db *sql.DB) RoomRepository {
	return &roomRepository{db: db}
}
