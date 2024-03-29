package session

import (
	"math/rand"
	"time"

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

func GenSessionToken() string {
	return RandomString(100)
}

const chars = "abcdefghijklmnopqrstuvwxyzabCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seed *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func RandomString(slen int) string {
	token := make([]byte, slen)
	for i := range token {
		token[i] = chars[seed.Intn(len(chars))]
	}
	return string(token)
}
