package config

const (
	SelectRoomFacilityList     = `SELECT id, room_id, facility_id, quantity, description, created_at, updated_at FROM trx_room_facility ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	SelectRoomFacilityByID     = `SELECT id, room_id, facility_id, quantity, description, created_at, updated_at FROM trx_room_facility WHERE id = $1`
	UpdateRoomFacility         = `UPDATE trx_room_facility SET room_id = $1, facility_id = $2, quantity = $3, description= $4, updated_at = CURRENT_TIMESTAMP WHERE id=$5 RETURNING created_at, updated_at`
	GetCountRoomFacility       = `SELECT COUNT(*) FROM trx_room_facility`
	GetQuantityFacilityByID    = `SELECT quantity FROM facilities WHERE id = $1`
	UpdateQuantityFacilityByID = `UPDATE facilities SET quantity = $1 WHERE id = $2`
	InsertTrxRoomFacility      = `INSERT INTO trx_room_facility (room_id, facility_id, quantity, description, updated_at) VALUES ($1, $2, $3, $4,CURRENT_TIMESTAMP) RETURNING id, created_at, updated_at`

	SelectTransactionList         = `SELECT id, employee_id, room_id, description, status, start_time, end_time, created_at, updated_at FROM transactions WHERE created_at BETWEEN $3 AND ($4::date + 1) - interval '1 second' ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	SelectRoomWithFacilities      = `SELECT r.id, r.facility_id, r.quantity, r.description, r.created_at, r.updated_at FROM rooms t JOIN trx_room_facility r on t.id = r.room_id WHERE t.id = $1;`
	GetIdListTransaction          = `SELECT COUNT(*) FROM transactions`
	GetEmployeeIdListTransaction  = `SELECT COUNT(*) FROM transactions WHERE employee_id = $1`
	SelectTransactionByID         = `SELECT id, employee_id, room_id, description, status, start_time, end_time, created_at, updated_at FROM transactions WHERE id = $1`
	SelectTransactionByEmployeeID = `SELECT id, employee_id, room_id, description, status, start_time, end_time, created_at, updated_at FROM transactions WHERE employee_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
	InsertTransactions            = `INSERT INTO transactions (employee_id, room_id, description, start_time, end_time, updated_at) VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP) RETURNING id, status, created_at, updated_at`
	UpdatePermission              = `UPDATE transactions SET status = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2 RETURNING employee_id, room_id, description, start_time, end_time, created_at`
	InsertRoomFacility            = `INSERT INTO trx_room_facility (room_id, facility_id, quantity, description, updated_at) VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP) RETURNING id, created_at, updated_at`
	UpdateFacilityQuantity        = `UPDATE facilities SET quantity = quantity - $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2 RETURNING id, created_at, updated_at`
	SelectQuantityFacility        = `SELECT quantity FROM facilities WHERE id = $1`
	SelectRoomByID2               = `SELECT status FROM rooms WHERE id = $1`
	// `SELECT id, date, amount, transaction_type, balance, description, created_at, updated_at FROM expenses WHERE LOWER(transaction_type::text) = LOWER($1)`

	InsertRoom            = `INSERT INTO rooms (name, room_type, capacity, status) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`
	SelectRoomByID        = `SELECT id, name, room_type, capacity, status, created_at, updated_at FROM rooms WHERE id = $1`
	SelectRoomList        = `SELECT id, name, room_type, capacity, status, created_at, updated_at FROM rooms ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	SelectRoomListStatus  = `SELECT id, name, room_type, capacity, status, created_at, updated_at FROM rooms WHERE status = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
	UpdateRoomByID        = `UPDATE rooms SET name = $2, room_type = $3, capacity = $4, status = $5, updated_at = CURRENT_TIMESTAMP WHERE id = $1 RETURNING created_at, updated_at`
	UpdateRoomStatus      = `UPDATE rooms SET status = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $1 RETURNING name, room_type, capacity, created_at, updated_at`
	SelectCountRoom       = `SELECT COUNT(*) FROM rooms`
	SelectCountRoomStatus = `SELECT COUNT(*) FROM rooms WHERE status = $1`

	InsertFasilities     = `INSERT INTO facilities (name, quantity) VALUES ($1, $2) RETURNING id, created_at, updated_at`
	SelectFasilitiesList = `SELECT id, name, quantity, created_at, updated_at FROM facilities ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	SelectFasilitiesById = `SELECT id, name, quantity, created_at, updated_at FROM facilities WHERE id = $1`
	UpdateFasilities     = `UPDATE facilities SET name = $1, quantity = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3 RETURNING created_at, updated_at`
	TotalRowsFasilities  = `SELECT COUNT(*) FROM facilities`

	// Employee
	// done
	InsertEmployee    = "INSERT INTO employees(name, username, password, role, division, position, contact, updated_at) VALUES($1, $2, crypt($3, gen_salt('bf')), $4, $5, $6, $7, CURRENT_TIMESTAMP) RETURNING id, created_at, updated_at;"
	SelectAllEmployee = "SELECT id, name, username, password, role, division, position, contact, created_at, updated_at FROM employees LIMIT $1 OFFSET $2;"
	// done
	SelectEmployeeByID       = "SELECT id, name, username, password, role, division, position, contact, created_at, updated_at FROM employees WHERE id = $1;"
	SelectEmployeeByUsername = "SELECT id, name, username, password, role, division, position, contact, created_at, updated_at FROM employees WHERE username = $1;"
	SelectEmployeeForLogin   = `SELECT id, name, username, password, role FROM employees WHERE username = $1 AND password = crypt($2, password)`
	// done

	UpdateEmployee = `UPDATE employees SET name = $1, username = $2, password = crypt($3, password), role = $4, division = $5, position = $6, contact = $7, updated_at = CURRENT_TIMESTAMP WHERE id = $8 RETURNING created_at, updated_at`

	SelectReportList             = `SELECT t.id, t.employee_id, e.name, e.username, e.division, e.position, e.contact, t.room_id, r.name, r.room_type, r.capacity, t.description, t.status, t.start_time, t.end_time, t.created_at, t.updated_at FROM transactions t JOIN employees e on e.id = t.employee_id JOIN rooms r on r.id = t.room_id WHERE t.created_at BETWEEN $1 AND $2 ORDER BY created_at DESC`
	SelectReportFacilityByRoomID = `SELECT t.facility_id, f.name, t.quantity FROM trx_room_facility t JOIN facilities f ON t.facility_id = f.id WHERE t.room_id = $1`
)
