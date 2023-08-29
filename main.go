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
	defer conn.Close()

	_, err = conn.Do("SET", "na", "new_value")
	if err != nil {
		fmt.Println(err)
	}

	value, err := redis.String(conn.Do("GET", "na"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n", value)
		return
	}

}
