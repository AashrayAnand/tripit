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
	Name string  `form:"name" json:"name" binding:"required"`
	X    float32 `form:"x" json:"x" binding:"required"`
	Y    float32 `form:"y" json:"y" binding:"required"`
}

type Trip struct {
	Id        uuid.UUID
	Locations []Location
}

type LocationList struct {
	Loc1 Location `form:"loc1" json:"loc1" binding:"required"`
}
