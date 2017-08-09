package main

import (
		"fmt"
		"unsafe"
		"reflect"
)

func main() {
		type T struct {
				i int
				j float32
		}
		var t T = T{1, 2.0}
		x := 2
		t2 := &T{2, 3.0}
		fmt.Println(t.i)
		fmt.Println(t2.j)
		fmt.Println(unsafe.Sizeof(T{}))
		fmt.Println(reflect.TypeOf(x))
}
