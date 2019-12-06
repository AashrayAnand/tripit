package main

import (
	"fmt"
	"log"
	"time"

	"github.com/AashrayAnand/tripit/secret"
	"github.com/AashrayAnand/tripit/session"
	"github.com/AashrayAnand/tripit/util"
	"googlemaps.github.io/maps"
)

func main() {
	// INIT GMAPS API  CLIENT
	_, err := maps.NewClient(maps.WithAPIKey(secret.GOOG_API_KEY))
	if err != nil {
		log.Fatalf("fatal error: %s", err.Error())
	}
	// TEST REDIS IS WORKING
	_, _ = session.Client.SetNX("x", "y", 60*time.Second).Result()
	val, err := session.Client.Get("x").Result()
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}
	fmt.Println(val)

	//router.Use(static.Serve("/", static.LocalFile("./views", true)))
	router := util.RunServer()
	router.Run(":3000") // listen and server 0.0.0.0:8080
}
