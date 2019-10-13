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
	Name  string  `form:"name" json:"name" binding:"required"`
	Xcoor float32 `form:"Xcoor" json:"Xcoor" binding:"required"`
	Ycoor float32 `form:"Ycoor" json:"Ycoor" binding:"required"`
}

type Trip struct {
	Id        uuid.UUID
	Locations []Location
}
