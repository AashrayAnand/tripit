package user

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AashrayAnand/tripit/driver"
	"github.com/AashrayAnand/tripit/models"
	"github.com/AashrayAnand/tripit/session"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

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
	// get user data from POST form
	user := c.PostForm("user")
	name := c.PostForm("name")
	pass := c.PostForm("pass")
	var searchResponse models.User
	// search for existing user with specified username
	err := driver.FindUser(user, &searchResponse)
	if err != nil {
		// user does not exist, create user
		if err.Error() == "mongo: no documents in result" {
			var newUser models.User
			createErr := createUserObject(&newUser, user, pass, name)
			if createErr != nil {
				res := fmt.Sprintf("error hashing password")
				c.JSON(301, gin.H{"message": res, "status": http.StatusInternalServerError})
				return
			}
			// TODO: add user with newUser struct, instead of passing fields
			addErr := driver.AddUser(newUser.Username, newUser.Password)
			if addErr != nil {
				res := fmt.Sprintf("error creating user")
				c.JSON(301, gin.H{"message": res, "status": http.StatusInternalServerError})
				return
			}
		} else {
			res := fmt.Sprintf("error finding user")
			c.JSON(301, gin.H{"message": res, "status": http.StatusInternalServerError})
			return
		}
	} else {
		res := fmt.Sprintf("user %s already exists", user)
		c.JSON(http.StatusBadRequest, gin.H{"message": res, "status": http.StatusBadRequest})
		return
	}

	res := fmt.Sprintf("user %s created", user)
	sessionId := AddSessionToken(user)
	c.JSON(http.StatusOK, gin.H{"message": res, "status": http.StatusOK, "token": sessionId})
}

// TODO: implement login, should authenticate user, and return
// appropriate response code, as well as a cookie, if the user
// was successfully authenticated
func Login(c *gin.Context) {
	// get user data from POST form
	user := c.PostForm("user")
	pass := c.PostForm("pass")

	var searchResponse models.User
	// search for existing user with specified username
	err := driver.FindUser(user, &searchResponse)
	if err != nil { // user does not exist
		if err.Error() == "mongo: no documents in result" {
			res := fmt.Sprintf("user does not exist")
			c.JSON(301, gin.H{"message": res, "status": http.StatusInternalServerError})
			return
		} else { // error querying db
			res := fmt.Sprintf("error finding user")
			c.JSON(301, gin.H{"message": res, "status": http.StatusInternalServerError})
			return
		}
	} else { // user exists, but specified passsword is incorrect
		if err := bcrypt.CompareHashAndPassword([]byte(searchResponse.Password), []byte(pass)); err != nil {
			res := fmt.Sprintf("incorrect user/pass combination %+v", searchResponse.Password)
			c.JSON(301, gin.H{"message": res, "status": http.StatusInternalServerError})

		} else { // user authenticated successfully
			// add session ID to session store for 30 minutes
			sessionId := AddSessionToken(user)
			res := fmt.Sprintf("welcome %s!", user)
			c.JSON(http.StatusOK, gin.H{"message": res, "status": http.StatusOK, "auth": sessionId})
		}
	}
}

func AddSessionToken(user string) string {
	sessionId := session.GenSessionToken()
	uuid := driver.GetId(user)
	// create session token, lasts 20 minutes
	_, err := session.Client.SetNX(sessionId, uuid, 60*20*time.Second).Result()
	if err != nil {
		log.Fatal("error on redis token add", err.Error())
	}
	return sessionId
}

// add all routes that are part of the /user group to the specified routing engine
func AddUserRoutes(router *gin.Engine) {
	users := router.Group("/user") // create user route groups /user/...
	{
		users.POST("/create", Create) // add create user route
		users.POST("/login", Login)   // add login route
		users.POST("/asd", func(c *gin.Context) {
			res := fmt.Sprintf("%#v", c.PostForm("trip"))
			c.JSON(http.StatusOK, gin.H{"message": res})
		})
	}
}

func createUserObject(newUser *models.User, user string, pass string, name string) error {
	newUser.Username = user
	newUser.Name = name
	// hash the user's pw
	hashPass, cryptErr := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	newUser.Password = string(hashPass)
	return cryptErr
}
