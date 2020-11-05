package main

import (
	"fmt"
)

var cache = make(map[int]int)
var lastVal int = 1

func fibo(num int) int {

	if num >= 0 {
		cache[0], cache[1] = 0, 1

		value, status := cache[num]
		if status == true {
			return value
		} else {

			for i := lastVal + 1; i <= num; i++ {
				cache[i] = cache[i-1] + cache[i-2]
			}
			lastVal = num
			return cache[num]
		}
	} else {
		return -1
	}
}

func main() {
	var input int
	fmt.Scanf("%v", &input)

	var result = fibo(input)

	if result >= 0 {
		fmt.Printf("Fibonacci number at %dth place is %d", input, result)
	} else {
		fmt.Printf("Fibonacci number for negative index is not available")
	}
}
