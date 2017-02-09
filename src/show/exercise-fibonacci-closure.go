package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	pre := 0
	now := 1
	return func() int {
		tmp := now
		now = now + pre
		pre = tmp
		return now
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}