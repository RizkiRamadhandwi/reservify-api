package repository

import (
	"booking-room-app/config"
	"booking-room-app/entity"
	"booking-room-app/shared/model"
	"database/sql"
	"log"
	"math"
)

type FasilitiesRepository interface {
	List(page, size int) ([]entity.Facilities, model.Paging, error)
	Create(payload entity.Facilities) (entity.Facilities, error)
	GetById(id string) (entity.Facilities, error)
	UpdateById(payload entity.Facilities) (entity.Facilities, error)
}

type fasilitiesRepository struct {
	db *sql.DB
}

// create facilities (ADMIN) -POST
// Create implements FasilitiesRepository.
func (f *fasilitiesRepository) Create(payload entity.Facilities) (entity.Facilities, error) {
	var fasilities entity.Facilities

	err := f.db.QueryRow(config.InsertFasilities,
		payload.Name,
		payload.Quantity).Scan(
		&fasilities.ID,
		&fasilities.CreatedAt,
		&fasilities.UpdatedAt)

	if err != nil {
		log.Println("fasilities repository.QueryRow:", err.Error())
		return entity.Facilities{}, err
	}
	fasilities.Name = payload.Name
	fasilities.Quantity = payload.Quantity
	fasilities.UpdatedAt = fasilities.CreatedAt
	return fasilities, nil
}

// List facilities (ADMIN) -POST
// GetAll implements FasilitiesRepository.
func (f *fasilitiesRepository) List(page, size int) ([]entity.Facilities, model.Paging, error) {
	var facilities []entity.Facilities
	offset := (page - 1) * size

	rows, err := f.db.Query(config.SelectFasilitiesList, size, offset)

	if err != nil {
		log.Println("fasilities repository.Query: ", err.Error())
		return nil, model.Paging{}, err
	}

	for rows.Next() {
		var facility entity.Facilities
		err := rows.Scan(
			&facility.ID,
			&facility.Name,
			&facility.Quantity,
			&facility.CreatedAt,
			&facility.UpdatedAt,
		)

		if err != nil {
			log.Println("scan facility : ", err.Error())
			return nil, model.Paging{}, err
		}

		facilities = append(facilities, facility)
	}

	totalRows := 0
	if err := f.db.QueryRow(config.TotalRowsFasilities).Scan(&totalRows); err != nil {
		return nil, model.Paging{}, err
	}

	paging := model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}

	return facilities, paging, nil
}

// get facility details by ID (ADMIN) -GET
// GetById implements FasilitiesRepository.
func (f *fasilitiesRepository) GetById(id string) (entity.Facilities, error) {
	var fasilities entity.Facilities
	err := f.db.QueryRow(config.SelectFasilitiesById, id).Scan(
		&fasilities.ID,
		&fasilities.Name,
		&fasilities.Quantity,
		&fasilities.CreatedAt,
		&fasilities.UpdatedAt)

	if err != nil {
		log.Println("fasilitiesRepository.Get.QueryRow:", err.Error())
		return entity.Facilities{}, err
	}

	return fasilities, nil

}

// Update facility details by ID (ADMIN) -PUT
// Update implements FasilitiesRepository.
func (f *fasilitiesRepository) UpdateById(payload entity.Facilities) (entity.Facilities, error) {
	var fasilities entity.Facilities

	err := f.db.QueryRow(config.UpdateFasilities,
		payload.Name,
		payload.Quantity,
		payload.ID).Scan(
		&fasilities.CreatedAt, &fasilities.UpdatedAt)

	if err != nil {
		log.Println("fasilitiesRepository.query:", err.Error())
		return entity.Facilities{}, err
	}

	fasilities.ID = payload.ID
	fasilities.Name = payload.Name
	fasilities.Quantity = payload.Quantity

	return fasilities, nil
}

func NewFasilitesRepository(db *sql.DB) FasilitiesRepository {
	return &fasilitiesRepository{db: db}
}
