package main

import "fmt"

func calc(x, y int) (int, int) {
	var res1 = x + y
	var res2 = x - y
	return res1, res2
}

/*method 2
func calc(x, y int) (res1,res2 int) {
	 res1 = x + y
	 res2 = x - y
	return
}*/

func main() {
	res1, res2 := calc(3, 1)
	fmt.Println(res1, res2)
}
