package trip

import (
	"net/http"

	"github.com/AashrayAnand/tripit/models"
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
	var list models.LocationList
	if err := c.ShouldBindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(301, gin.H{"message": list.Loc1, "status": http.StatusOK})
	return
}
