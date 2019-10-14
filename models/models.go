package models

import (
	"time"
)

// structure to represent user, will be populated
// upon new user creation
type User struct {
	Id       string // universally unique identifier
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
	Id        string
	Locations LocationList
}

type LocationList struct {
	Auth      string     `form:"auth" json:"auth" binding:"required"`
	Locations []Location `form:"locations" json:"locations" binding:"required"`
}
