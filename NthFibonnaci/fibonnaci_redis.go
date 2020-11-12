package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func fibo(num int) int {

	if num >= 0 {
		conn := pool.Get()
		length, err := redis.Int(conn.Do("HLEN", "Fibonnaci"))
		if err != nil {
			log.Fatal(err)
		}
		for i := length; i <= num; i++ {
			//	cache[i] = cache[i-1] + cache[i-2]
			num1, err := redis.Int(conn.Do("HGET", "Fibonnaci", i-1))
			num2, err := redis.Int(conn.Do("HGET", "Fibonnaci", i-2))
			num3 := num1 + num2
			_, err = conn.Do("HSETNX", "Fibonnaci", i, num3)
			if err != nil {
				log.Fatal(err)
			}
		}

		value, err := redis.Int(conn.Do("HGET", "Fibonnaci", num))
		if err != nil {
			log.Fatal(err)
		}

		return value //cache[num]
	} else {
		return -1
	}
}
func main() {
	pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
	conn := pool.Get()

	defer conn.Close()

	conn.Do("HSETNX", "Fibonnaci", 0, 0)

	conn.Do("HSETNX", "Fibonnaci", 1, 1)

	var input, result int
	fmt.Scanf("%v", &input)
	if num, err := redis.Int(conn.Do("HGET", "Fibonnaci", input)); err == nil {
		result = num
	} else {
		result = fibo(input)
		//fmt.Printf("notpresent")
	}

	if result >= 0 {
		fmt.Printf("Fibonacci number at %dth place is %d", input, result)
	} else {
		fmt.Printf("Fibonacci number for negative index is not available")
	}
}
