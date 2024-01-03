package repository

import (
	"booking-room-app/config"
	"booking-room-app/entity/dto"
	"database/sql"
	"log"
	"time"
)

type ReportRepository interface {
	List(startDate, endDate time.Time) ([]dto.ReportDto, error)
}

type reportRepository struct {
	db *sql.DB
}

// List implements ReportRepository.
func (r *reportRepository) List(startDate, endDate time.Time) ([]dto.ReportDto, error) {
	var reports []dto.ReportDto

	rows, err := r.db.Query(config.SelectReportList, startDate, endDate)
	if err != nil {
		log.Println("transactionsRepository.Query:", err.Error())
		return nil, err
	}
	for rows.Next() {
		var report dto.ReportDto
		err = rows.Scan(
			&report.ID,
			&report.EmployeeId,
			&report.Employee.Name,
			&report.Employee.Username,
			&report.Employee.Division,
			&report.Employee.Position,
			&report.Employee.Contact,
			&report.RoomId,
			&report.Room.Name,
			&report.Room.RoomType,
			&report.Room.Capacity,
			&report.Description,
			&report.Status,
			&report.StartTime,
			&report.EndTime,
			&report.CreatedAt,
			&report.UpdatedAt)
		if err != nil {
			log.Println("transactionsRepository.Rows.Next():", err.Error())
			return nil, err
		}

		roomFacilities, err := r.db.Query(config.SelectReportFacilityByRoomID, report.RoomId)
		if err != nil {
			log.Println("transactionsRepository.Query:", err.Error())
			return nil, err
		}
		for roomFacilities.Next() {
			var roomFacility dto.RoomFacilityDto
			err = roomFacilities.Scan(
				&roomFacility.FacilityID,
				&roomFacility.Name,
				&roomFacility.Quantity)
			if err != nil {
				log.Println("transactionsRepository.Rows.Next():", err.Error())
				return nil, err
			}
			report.RoomFacilities = append(report.RoomFacilities, roomFacility)
		}

		reports = append(reports, report)
	}

	return reports, nil
}

func NewReportRepository(db *sql.DB) ReportRepository {
	return &reportRepository{db: db}
}
