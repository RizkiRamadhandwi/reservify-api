CREATE DATABASE booking_room_db;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TYPE role_type AS ENUM ('employee', 'admin', 'ga');

CREATE TABLE employees (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(200) NOT NULL,
    division VARCHAR(50) NOT NULL,
    position VARCHAR(50) NOT NULL,
    role role_type DEFAULT 'employee',
    contact VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);


CREATE TABLE facilities (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name     VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    quantity INT NOT NULL
);

CREATE TYPE status_type AS ENUM ('available', 'booked', 'unavailable' );

CREATE TABLE rooms (
    id   uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name      VARCHAR(100) NOT NULL,
    room_type VARCHAR(100) NOT NULL,
    capacity  INT NOT NULL,
    status status_type DEFAULT 'available', 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE trx_room_facility (
    id uuid DEFAULT uuid_generate_v4() UNIQUE,
    room_id         uuid NOT NULL,
    facility_id     uuid NOT NULL,
    quantity        INT NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (room_id) REFERENCES rooms(id),
    FOREIGN KEY (facility_id) REFERENCES facilities(id)
);

CREATE TYPE transaction_status AS ENUM ('pending', 'accepted', 'declined');

CREATE TABLE transactions (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    employee_id uuid NOT NULL,
    room_id uuid NOT NULL,
    description TEXT,
    status transaction_status DEFAULT 'pending',
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (employee_id) REFERENCES employees(id),
    FOREIGN KEY (room_id) REFERENCES rooms(id)
);