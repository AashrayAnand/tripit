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
	Auth string   `form:"auth" json:"auth" binding:"required"`
	Loc1 Location `form:"loc1" json:"loc1" binding:"required"`
	Loc2 Location `form:"loc2" json:"loc2" binding:"required"`
	Loc3 Location `form:"loc3" json:"loc3" binding:"required"`
	Loc4 Location `form:"loc4" json:"loc4" binding:"required"`
	Loc5 Location `form:"loc5" json:"loc5" binding:"required"`
}
