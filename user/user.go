package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// structure to represent user, will be populated
// upon new user creation
type User struct {
	Username string `json:user`
	Name     string `json:name`
	Password string `json:pass`
	Token    string `json:token`
}

// user routing group should include all necessary functionality to

// 1. create new users
// 2. authenticate existing users

// Content-Type: application/x-www-form-urlencoded
// POST form parameters will be in following form
// user=<username>&pass=<pass>
// check if querying users collection for given user
// name returns an existing user, if not, hash the
// password and create a new user in users collection
func Create(c *gin.Context) {
	// gin.H is short for map[string]interface{}
	user := c.PostForm("user")
	pass := c.PostForm("pass")
	res := fmt.Sprintf("user %s with pass %s created", user, pass)
	c.JSON(200, gin.H{"message": res, "status": http.StatusOK})
}

// TODO: implement login
func Login(c *gin.Context) {
	c.JSON(200, nil)
}

// add all routes that are part of the /user group to the specified routing engine
func AddUserRoutes(router *gin.Engine) {
	users := router.Group("/user") // create user route groups /user/...
	{
		users.POST("/create", Create) // add create user route
		users.POST("/login", Login)   // add login route
	}
}
