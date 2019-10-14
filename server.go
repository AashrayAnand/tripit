package main

import (
	"fmt"
	"time"

	"github.com/AashrayAnand/tripit/session"
	"github.com/AashrayAnand/tripit/util"
)

func main() {

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
