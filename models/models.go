package models

import (
	"time"

	"github.com/google/uuid"
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
