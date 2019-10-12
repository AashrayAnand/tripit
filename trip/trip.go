package trip

import {
	"fmt"
	"net/http"

	"github.com/AashrayAnand/tripit/driver"
	"github.com/AashrayAnand/tripit/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
}

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
	// get trip data from POST form
	trip := c.PostForm("trip") // [{}, {}, ...]
	res := fmt.Print(trip)
	c.JSON(301, gin.H{"message": res, "status": http.OK})
	return
}
