package models

import (
	"time"

	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

// structure to represent user, will be populated
// upon new user creation
type User struct {
	Id       uuid.UUID // universally unique identifier
	Created  time.Time
	Username string
	Name     string
	Password string
}

type Location struct {
	Name	 string
	X-coor	 float32
	Y-coor	 float32
}

type Trip struct {
	Id		  uuid.UUID
	Locations []Location{}
}