package main

import (
	"github.com/gomodule/redigo/redis"
	"os"
)

var (
	Pool *redis.Pool
)

func Incr(counterKey string) error {

	conn, err := redis.Dial("tcp", os.Getenv("REDIS_HOST"))
	_, err = redis.Int(conn.Do("INCR", counterKey))

	conn.Close()
	return err
}