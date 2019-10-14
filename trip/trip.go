package trip

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AashrayAnand/tripit/driver"
	"github.com/AashrayAnand/tripit/models"
	"github.com/AashrayAnand/tripit/session"
	"github.com/gin-gonic/gin"
)

// trip routing group should include all necessary functionality to

// 1. create new trip
// 2. add a location to a trip
// 3. delete a trip

// Content-Type: json
// POST form parameters will be in following form
// user=<username>&pass=<pass>
// check if querying trip collection for a given user
// name returns an existing trip, if not, return "no trips"
func Create(c *gin.Context) {

	// bind request data
	var list models.LocationList
	if err := c.ShouldBindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if list.Auth == "" {
		c.JSON(400, gin.H{"message": "unauthorized", "status": http.StatusBadRequest})
		return
	}

	// get user's UUID from redis, using session token
	// add trip to mongodb, check for error
	uuid, err := session.Client.Get(list.Auth).Result()
	if err != nil {
		log.Fatal(err.Error())
	}

	driver.AddTrip(list, uuid)

	res := fmt.Sprintf("uuid is %s", uuid)
	c.JSON(301, gin.H{"message": res, "status": http.StatusOK})
	return
}
