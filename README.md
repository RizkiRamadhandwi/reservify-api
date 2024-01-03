# Resevify Aplication

## Prerequisites

Before running the Resevify application, make sure you have fulfilled the following prerequisites:

- Go (Golang) is installed on your system.
- PostgreSQL is installed, and you have created the tables as specified in the `ddl.sql` file. Then, insert the table contents from the `dml.sql` file as dummy data.
- An active internet connection is required to download Go dependencies.

## Running the Application

Once the application is running, you can access it through a web browser or use it through an API client such as Postman or cURL. Then, you can log in using an account created by the admin. This application provides APIs for managing Rooms, Facilities, Employees, and Transactions.

## Using the API

Below are instructions on how to use the API based on the features provided by the Resevify application:

### API Spec

#### Login API {Admin, Employee, GA}

Request :

- Method : `POST`
- Endpoint : `/employees`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
  "username": "string",
  "password": "string"
}
```

#### Employee API

##### Create Employee {Admin}

Request :

- Method : POST
- Endpoint : `/employees`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token
- Body :

```json
{
  "name": "string",
  "username": "string",
  "password": "string",
  "role": "string",
  "division": "string",
  "position": "string",
  "contact": "string"
}
```

Response :

- Status : 201 Created
- Body :

```json
{
  "status": {
    "code": 201,
    "message": "Created"
  },
  "data": {
    "id": "string",
    "name": "string",
    "username": "string",
    "password": "string",
    "role": "string",
    "division": "string",
    "position": "string",
    "contact": "string",
    "createdAt": "2000-01-01T12:00:00Z", (curent time)
    "updatedAt": "2000-01-01T12:00:00Z"  (curent time)
  }
}
```

##### Get Employees {Admin, Employee, GA}

Request :

- Method : GET
- Endpoint : `/employees`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Query Param :
  - page : int `optional`
  - size : int `optional`
- Authorization : Bearer Token

Response :

- Status : 200 OK
- Body :

```json
{
    "status": {
        "code": 200,
        "message": "Ok"
    },
    "data": [
        {
            "id": "string",
            "name": "string",
            "username": "string",
            "password": "string",
            "role": "string",
            "division": "string",
            "position": "string",
            "contact": "string",
            "createdAt": "2000-01-01T12:00:00Z",
            "updatedAt": "2000-01-01T12:00:00Z"
        }
    ],
    "paging": {
        "page": 1,          (default value)
        "rowsPerPage": 5,   (default value)
        "totalRows": int,
        "totalPages": int
    }
}

```

##### Get Employee By Id {Admin, Employee, GA}

Request :

- Method : GET
- Endpoint : `/employees/:id`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token

Response :

- Status : 200 OK
- Body :

```json
{
  "status": {
    "code": 200,
    "message": "Ok"
  },
  "data": {
    "id": "string",
    "name": "string",
    "username": "string",
    "password": "string",
    "role": "string",
    "division": "string",
    "position": "string",
    "contact": "string",
    "createdAt": "2000-01-01T00:00:00Z",
    "updatedAt": "2000-01-01T00:00:00Z"
  }
}
```

##### Get By Employee Usename {Admin, Employee, GA}

Request :

- Method : GET
- Endpoint : `/employees/username/:username`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token

Response :

- Status : 200 OK
- Body :

```json
{
  "status": {
    "code": 200,
    "message": "Ok"
  },
  "data": {
    "id": "string",
    "name": "string",
    "username": "string",
    "password": "string",
    "role": "string",
    "division": "string",
    "position": "string",
    "contact": "string",
    "createdAt": "2000-01-01T00:00:00Z",
    "updatedAt": "2000-01-01T00:00:00Z"
  }
}
```

##### Update Employee {Admin}

Request :

- Method : PUT
- Endpoint : `/employees`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token

```json
{
  "id": "string",
  "name": "string",
  "username": "string",
  "password": "string",
  "role": "string",
  "division": "string",
  "position": "string",
  "contact": "string"
}
```

Response :

- Status : 200 OK
- Body :

```json
{
  "status": {
    "code": 200,
    "message": "Updated Successfully"
  },
  "data": {
    "id": "string",
    "name": "string",
    "username": "string",
    "password": "string",
    "role": "string",
    "division": "string",
    "position": "string",
    "contact": "string",
    "createdAt": "2000-01-01T00:00:00Z",
    "updatedAt": "2000-01-01T00:00:00Z" (current time)
  }
}
```

#### Facility API

##### Create Facility {Admin}

Request :

- Method : POST
- Endpoint : `/facilities`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token
- Body :

```json
{
    "name": "string",
    "quantity": int
}
```

Response :

- Status : 201 Created
- Body :

```json
{
    "status": {
        "code": 201,
        "message": "Created"
    },
    "data": {
        "id": "string",
        "name": "string",
        "quantity": int,
        "createdAt": "2000-01-01T00:00:00Z", (current time)
        "updatedAt": "2000-01-01T00:00:00Z" (current time)
    }
}
```

##### Get Facilities {Admin, Employee, GA}

Request :

- Method : GET
- Endpoint : `/facilities`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Query Param :
  - page : int `optional`
  - size : int `optional`

Response :

- Status : 200 OK
- Body :

```json
{
    "status": {
        "code": 200,
        "message": "ok"
    },
    "data": [
        {
            "id": "string",
            "name": "string",
            "quantity": int,
            "createdAt": "2000-01-01T00:00:00Z",
            "updatedAt": "2000-01-01T00:00:00Z"
        }
    ],
    "paging": {
        "page": 1, (default value)
        "rowsPerPage": 5, (default value)
        "totalRows": int,
        "totalPages": int
    }
}
```

##### Get Facility By Id {Admin, Employee, GA}

Request :

- Method : GET
- Endpoint : `/facilities/:id`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token

Response :

- Status : 200 OK
- Body :

```json
{
    "status": {
        "code": 200,
        "message": "ok"
    },
    "data": {
        "id": "string",
        "name": "string",
        "quantity": int,
        "createdAt": "2000-01-01T00:00:00Z",
        "updatedAt": "2000-01-01T00:00:00Z"
    }
}

```

##### Update Facility By Id {Admin}

Request :

- Method : GET
- Endpoint : `/facilities/:id`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token

```json
{
    "id": "string",
    "name": "string",
    "quantity": int
}
```

Response :

- Status : 200 OK
- Body :

```json
{
    "status": {
        "code": 200,
        "message": "Updated"
    },
    "data": {
        "id": "string",
        "name": "string",
        "quantity": int,
        "createdAt": "2000-01-01T00:00:00Z",
        "updatedAt": "2000-01-01T00:00:00Z" (curent time)
    }
}

```

#### Room API

##### Create Room {Admin}

Request :

- Method : POST
- Endpoint : `/rooms`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token
- Body :

```json
{
    "name": "string",
    "room_type": "string",
    "capacity": int,
    "status": "string"
}
```

Response :

- Status : 201 Created
- Body :

```json
{
    "status": {
        "code": 201,
        "message": "Created"
    },
    "data": {
        "id": "string",
        "name": "string",
        "room_type": "string",
        "capacity": int,
        "status": "string",
        "createdAt": "2000-01-01T12:00:00Z", (curent time)
        "updatedAt": "2000-01-01T12:00:00Z"  (curent time)
    }
}
```

###### Get Rooms

Request :

- Method : GET
- Endpoint : `/rooms`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token
- Query Param :
  - page : int `optional`
  - size : int `optional`

Response :

- Status : 201 Ok
- Body :

```json
{
    "status": {
        "code": 200,
        "message": "Ok"
    },
    "data": [
        {
            "id": "string",
            "name": "string",
            "room_type": "string",
            "capacity": int,
            "status": "string",
            "created_at": "2000-01-01T00:00:00Z",
            "updated_at": "2000-01-01T00:00:00Z"
        }
    ],
    "paging": {
        "page": 1, (default value)
        "rowsPerPage": 5, (default value)
        "totalRows": int,
        "totalPages": int
    }
}
```

##### Get Room By Id {Admin, Employee, GA}

Request :

- Method : GET
- Endpoint : `/rooms/:id`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token
- Query Param :
  - page : int `optional`
  - size : int `optional`

Response :

- Status : 201 Ok
- Body :

```json
{
    "status": {
        "code": 200,
        "message": "Ok"
    },
    "data":{
        "id": "string",
        "name": "string",
        "room_type": "string",
        "capacity": int,
        "status": "string",
        "created_at": "2000-01-01T00:00:00Z",
        "updated_at": "2000-01-01T00:00:00Z"
    }
}
```

##### Get Room {Admin, Employee, GA}

Request :

- Method : GET
- Endpoint : `/rooms`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token
- Query Param :
  - page : int `optional`
  - size : int `optional`

Response :

- Status : 201 Ok
- Body :

```json
{
    "status": {
        "code": 200,
        "message": "Ok"
    },
    "data":{
        "id": "string",
        "name": "string",
        "room_type": "string",
        "capacity": int,
        "status": "string",
        "created_at": "2000-01-01T00:00:00Z",
        "updated_at": "2000-01-01T00:00:00Z"
    }
}
```

##### Update Rooms {Admin}

Request :

- Method : GET
- Endpoint : `/rooms`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token

```json
{
    "id": "string",
    "name": "string",
    "room_type": "string",
    "capacity": int,  (optional)
    "status": "string"  (optional)
}
```

Response :

- Status : 200 OK
- Body :

```json
{
    "status": {
        "code": 201,
        "message": "Updated"
    },
    "data": {
        "id": "string",
        "name": "string",
        "room_type": "string",
        "capacity": int,
        "status": "string",
        "created_at": "2000-01-01T00:00:00Z",
        "updated_at": "2000-01-01T00:00:00Z" (current time)
    }
}
```

##### Update Rooms Status {Admin}

Request :

- Method : GET
- Endpoint : `/rooms`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token

```json
{
  "id": "string",
  "status": "string"
}
```

Response :

- Status : 200 OK
- Body :

```json
{
    "status": {
        "code": 201,
        "message": "Updated"
    },
    "data": {
        "id": "string",
        "name": "string",
        "room_type": "string",
        "capacity": int,
        "status": "string",
        "created_at": "2000-01-01T00:00:00Z",
        "updated_at": "2000-01-01T00:00:00Z" (current time)
    }
}
```

#### Room Facility API

##### Create Room Facility {Admin}

Request :

- Method : POST
- Endpoint : `/roomfacilities`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token
- Body :

```json
{
    "roomId": "string",
    "facilityId": "string",
    "quantity": int,
    "description": "string"
}
```

Response :

- Status : 201 Created
- Body :

```json
{
  "status": {
      "code": 201,
      "message": "Created"
  },
  "data": {
      "id": "string",
      "roomId": "string",
      "facilityId": "string",
      "quantity": int,
      "description": "string",
      "createdAt": "2000-01-01T12:00:00Z", (curent time)
      "updatedAt": "2000-01-01T12:00:00Z"  (curent time)
  }
}
```

##### Get Room Facilities {Admin}

Request :

- Method : GET
- Endpoint : `/roomfacilities`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Query Param :
  - page : int `optional`
  - size : int `optional`
- Authorization : Bearer Token

Response :

- Status : 200 OK
- Body :

```json
{
    "status": {
        "code": 200,
        "message": "Ok"
    },
    "data": [
        {
            "id": "string",
            "roomId": "string",
            "facilityId": "string",
            "quantity": 5,
            "description": "string",
            "createdAt": "2000-01-01T12:00:00Z",
            "updatedAt": "2000-01-01T12:00:00Z"
        }
    ],
     "paging": {
        "page": 1,          (default value)
        "rowsPerPage": 5,   (default value)
        "totalRows": int,
        "totalPages": int
    }
}
```

##### Get Room Facility By Id {Admin}

Request :

- Method : GET
- Endpoint : `/roomfacilities/:id`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token

Response :

- Status : 200 OK
- Body :

```json
{
  "status": {
    "code": 200,
    "message": "Ok"
  },
  "data": {
    "id": "string",
    "roomId": "string",
    "facilityId": "string",
    "quantity": int,
    "description": "string",
    "createdAt": "2000-01-01T00:00:00Z",
    "updatedAt": "2000-01-01T00:00:00Z"
  }
}
```

##### Update Room Facility {Admin}

Request :

- Method : PUT
- Endpoint : `/roomfacilities`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token

```json
{
    "id": "string",
    "quantity": int,
    "description": "string"
}
```

Response :

- Status : 200 OK
- Body :

```json
{
    "status": {
        "code": 201,
        "message": "Updated"
    },
    "data": {
        "id": "string",
        "roomId": "string",
        "facilityId": "string",
        "quantity": int,
        "description": "string",
        "createdAt": "2000-01-01T00:00:00Z",
        "updatedAt": "2000-01-01T00:00:00Z" (current time)
    }
}
```

#### Transaction API

##### Create Transaction {Admin, Employee}

Request :

- Method : POST
- Endpoint : `/transactions`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token
- Body :

```json
{
        "employeeId": "string",
        "roomId": "string",
        "roomFacilities": [ (optional)
            {
                "facilityId": "string",
                "quantity": int,
                "description": "string"
            }
        ],
        "description": "string",
        "startTime": "2000-01-01T00:00:00Z",
        "endTime": "2000-01-01T01:00:00Z"
}
```

Response :

- Status : 201 Created
- Body :

```json
{
    "status": {
        "code": 201,
        "message": "Created"
    },
    "data": {
        "id": "string",
        "employeeId": "string",
        "roomId": "string",
        "roomFacilities": [
            {
                "id": "string",
                "facilityId": "string",
                "quantity": int,
                "description": "string",
                "createdAt": "2000-01-01T00:00:00Z", (current time)
                "updatedAt": "2000-01-01T00:00:00Z" (current time)
            }
        ],
        "description": "string",
        "status": "string",
        "startTime": "2000-01-01T00:00:00Z",
        "endTime": "2000-01-01T01:00:00Z",
        "createdAt": "2000-01-01T00:00:00Z", (current time)
        "updatedAt": "2000-01-01T00:00:00Z" (current time)
    }
}
```

##### Get Transactions {Admin, GA}

Request :

- Method : GET
- Endpoint : `/transactions`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Query Param :
  - page : int `optional`
  - size : int `optional`
  - startDate : date(yyyy-mm-dd) `optional`
  - endDate : date(yyyy-mm-dd) `optional`
- Authorization : Bearer Token

```json
{
    "status": {
        "code": 200,
        "message": "Ok"
    },
    "data": [
        {
            "id": "string",
            "employeeId": "string",
            "roomId": "string",
            "roomFacilities": [
                {
                    "id": "string",
                    "facilityId": "string",
                    "quantity": int,
                    "description": "string",
                    "createdAt": "2000-01-01T12:00:00Z",
                    "updatedAt": "2000-01-01T12:00:00Z"
                }
            ],
            "description": "string",
            "status": "string",
            "startTime": "2000-01-01T12:00:00Z",
            "endTime": "2000-01-01T12:00:00Z",
             "createdAt": "2000-01-01T12:00:00Z",
            "updatedAt": "2000-01-01T12:00:00Z"
        }
    ],
    "paging": {
        "page": 1,          (default value)
        "rowsPerPage": 5,   (default value)
        "totalRows": int,
        "totalPages": int
    }
}
```

##### Get Transaction By Id {Admin, Employee, GA}

Request :

- Method : GET
- Endpoint : `/transactions/:id`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token

Response :

- Status : 200 OK
- Body :

```json
{
    "status": {
        "code": 200,
        "message": "Ok"
    },
    "data": {
        "id": "string",
        "employeeId": "string",
        "roomId": "string",
        "roomFacilities": [
            {
                "id": "string",
                "facilityId": "string",
                "quantity": int,
                "description": "string",
                "createdAt": "2000-01-01T00:00:00Z",
                "updatedAt": "2000-01-01T00:00:00Z"
            }
        ],
        "description": "string",
        "status": "string",
        "startTime": "2000-01-01T00:00:00Z",
        "endTime": "2000-01-01T00:00:00Z",
        "createdAt": "2000-01-01T00:00:00Z",
        "updatedAt": "2000-01-01T00:00:00Z"
    }
}
```

##### Get Transaction By Employee Id {Admin, Employee, GA}

Request :

- Method : GET
- Endpoint : `/transactions/employee/:id`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token

Response :

- Status : 200 OK
- Body :

```json
{
    "status": {
        "code": 200,
        "message": "Ok"
    },
    "data": {
        "id": "string",
        "employeeId": "string",
        "roomId": "string",
        "roomFacilities": [
            {
                "id": "string",
                "facilityId": "string",
                "quantity": int,
                "description": "string",
                "createdAt": "2000-01-01T00:00:00Z",
                "updatedAt": "2000-01-01T00:00:00Z"
            }
        ],
        "description": "string",
        "status": "string",
        "startTime": "2000-01-01T00:00:00Z",
        "endTime": "2000-01-01T00:00:00Z",
        "createdAt": "2000-01-01T00:00:00Z",
        "updatedAt": "2000-01-01T00:00:00Z"
    }
}
```

##### Update Room Facility {Admin}

Request :

- Method : PUT
- Endpoint : `/transactions`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token

```json
{
    "status": {
        "code": 201,
        "message": "Updated"
    },
    "data": {
        "id": "string",
        "employeeId": "string",
        "roomId": "string",
        "description": "string",
        "status": "string",
        "startTime": "2000-01-01T00:00:00Z",
        "endTime": "2000-01-01T01:00:00Z",
        "createdAt": "2000-01-01T00:00:00Z", (current time)
        "updatedAt": "2000-01-01T00:00:00Z" (current time)
    }
}
```

#### Report API

##### Download Report {Admin}

Request :

- Method : GET
- Endpoint : `/reports/download`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token
- Query Param :
- range : string
