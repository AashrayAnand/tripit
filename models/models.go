package models

import "time"

// structure to represent user, will be populated
// upon new user creation
type User struct {
	Username string `json:user`
	Name     string `json:name`
	Password string `json:pass`
	Token    string `json:token`
	Created  time.Time
}
