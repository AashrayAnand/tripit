package trip

import (
	"fmt"
	"net/http"

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
	// get trip data from POST form
	trip := c.PostForm("trip") // [{}, {}, ...]
	res := fmt.Sprintf("%v", trip)
	c.JSON(301, gin.H{"message": res, "status": http.StatusOK})
	return
}
