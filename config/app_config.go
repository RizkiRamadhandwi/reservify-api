package config

const (
	ApiGroup = "/api/v1"

	// Rooms
	RoomCreate       = "/rooms"
	RoomList         = "/rooms"
	RoomGetById      = "/rooms/:id"
	RoomUpdateStatus = "/rooms/status"
	RoomUpdate       = "/rooms"
	// RoomDelete       = "/rooms/:id"

	// Facilities
	FacilitiesCreate  = "/facilities"
	FacilitiesList    = "/facilities"
	FacilitiesGetById = "/facilities/:id"
	FacilitiesUpdate  = "/facilities"

	// Employees
	EmployeesList    = "/employees"
	EmployeesCreate  = "/employees"
	EmployeesGetById = "/employees/:id"
	EmployeesGetByUsername = "/employees/username/:user"
	EmployeesUpdate  = "/employees"
	EmployeesDelete  = "/employees/:id"

	// Transaction
	TransactionList       = "/transactions"
	TransactionCreate     = "/transactions"
	TransactionGetById    = "/transactions/:id"
	TransactionGetByEmpId = "/transactions/employee/:employeeId"
	// TransactionPermList   = "/transactions"
	TransactionUpdatePerm = "/transactions/status"

	// Room Facilities
	RoomFacilityCreate  = "/roomfacilities"
	RoomFacilityList    = "/roomfacilities"
	RoomFacilityGetById = "/roomfacilities/:id"
	RoomFacilityUpdate  = "/roomfacilities"

	// Auth
	AuthLogin = "/auth/login"
)
