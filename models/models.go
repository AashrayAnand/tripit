package models

import (
	"time"
)

// structure to represent user, will be populated
// upon new user creation
type User struct {
	Id       string
	Username string
	Name     string
	Password string
	Token    string
	Created  time.Time
}
