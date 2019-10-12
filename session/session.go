package session

import (
	"github.com/go-redis/redis"
)

// initialize redis session store client
var Client = initSessionStore()

func initSessionStore() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
