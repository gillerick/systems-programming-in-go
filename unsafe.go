package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	var value int64 = 5

	var p1 = &value //Pointer to value
	var p2 = (*int64)(unsafe.Pointer(p1))


	fmt.Println("*p1: ", *p1)
	fmt.Println("*p2: ", *p2)
	*p1 = 3121213232212213212

	fmt.Println(value)
	fmt.Println("*p2: ", *p2)
	*p1 = 31212132

}
