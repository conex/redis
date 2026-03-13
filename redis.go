package redis

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/omeid/conex"
)

var (
	// Image to use for the box.
	Image = "redis:alpine"
	// Port used for connect to redis.
	Port = "6379"

	// RedisUpWaitTime dictates how long we should wait for Redis to accept connections.
	RedisUpWaitTime = 10 * time.Second
)

func init() {
	conex.Require(func() string { return Image })
}

// Box returns a redis.Client and the container running the redis
// server. It calls t.Fatal on errors.
func Box(t testing.TB, db int) (*redis.Client, conex.Container) {
	c := conex.Box(t, &conex.Config{
		Image:  Image,
		Expose: []string{Port},
	})

	t.Log("Waiting for Redis to accept connections")

	err := c.Wait(Port, RedisUpWaitTime)
	if err != nil {
		c.Drop()
		t.Fatal("Redis failed to start:", err)
	}

	t.Log("Redis is now accepting connections")

	addr := fmt.Sprintf("%s:%s", c.Address(), Port)
	opt := &redis.Options{
		Addr: addr,
		DB:   db,
	}

	client := redis.NewClient(opt)

	return client, c
}
