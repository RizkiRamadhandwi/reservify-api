package delivery

import (
	"booking-room-app/config"
	"booking-room-app/delivery/controller"
	"booking-room-app/delivery/middleware"
	"booking-room-app/repository"
	"booking-room-app/shared/service"
	"booking-room-app/usecase"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	roomUC         usecase.RoomUseCase
	facilitiesUC   usecase.FacilitiesUseCase
	employeeUC     usecase.EmployeesUseCase
	roomFacilityUc usecase.RoomFacilityUsecase
	transactionsUc usecase.TransactionsUsecase
	reportUC       usecase.ReportUseCase
	authUsc        usecase.AuthUseCase
	engine         *gin.Engine
	jwtService     service.JwtService
	host           string
}

func (s *Server) initRoute() {
	rg := s.engine.Group(config.ApiGroup)

	authMiddleware := middleware.NewAuthMiddleware(s.jwtService)
	controller.NewRoomController(s.roomUC, authMiddleware, rg).Route()
	controller.NewFacilitiesController(s.facilitiesUC, rg, authMiddleware).Route()
	controller.NewEmployeeController(s.employeeUC, rg, authMiddleware).Route()
	controller.NewRoomFacilityController(s.roomFacilityUc, rg, authMiddleware).Route()
	controller.NewTransactionsController(s.transactionsUc, rg, authMiddleware).Route()
	controller.NewAuthController(s.authUsc, rg).Route()
	controller.NewReportController(s.reportUC, rg, authMiddleware).Route()
}

func (s *Server) Run() {
	s.initRoute()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("server not running on host %s, becauce error %v", s.host, err.Error()))
	}
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		panic(err.Error())
	}

	// Inject DB ke -> repository
	roomRepo := repository.NewRoomRepository(db)
	facilityRepo := repository.NewFasilitesRepository(db)
	employeeRepo := repository.NewEmployeeRepository(db)
	roomFacilityRepo := repository.NewRoomFacilityRepository(db)
	transactionsRepo := repository.NewTransactionsRepository(db)
	reportRepo := repository.NewReportRepository(db)

	// Inject REPO ke -> useCase
	roomUC := usecase.NewRoomUseCase(roomRepo)
	facilitiesUC := usecase.NewFacilitiesUseCase(facilityRepo)
	employeeUC := usecase.NewEmployeeUseCase(employeeRepo)
	roomFacilityUc := usecase.NewRoomFacilityUsecase(roomFacilityRepo)
	transactionsUc := usecase.NewTransactionsUsecase(transactionsRepo)
	jwtService := service.NewJwtService(cfg.TokenConfig)
	authUc := usecase.NewAuthUseCase(employeeUC, jwtService)
	reportUC := usecase.NewReportUseCase(reportRepo)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		authUsc:        authUc,
		roomUC:         roomUC,
		facilitiesUC:   facilitiesUC,
		employeeUC:     employeeUC,
		transactionsUc: transactionsUc,
		roomFacilityUc: roomFacilityUc,
		reportUC:       reportUC,
		engine:         engine,
		jwtService:     jwtService,
		host:           host,
	}
}
