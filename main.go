package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")

	if err != nil {
		fmt.Println("Bad!")
	} else {
		fmt.Println("Good!!!")
	}

	_, err = conn.Do("SET", "na", "new_value")
	value, err := redis.String(conn.Do("MGET", "na", "key"))
	fmt.Println(value)
}
