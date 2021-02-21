package main

import (
	"fmt"
	"os"
	"strconv"
)

func main()  {
	arguments := os.Args
	sum := 0
	for i := 0; i < len(arguments) ; i++{
		temp, _ := strconv.Atoi(arguments[i])
		sum += temp
	}

	fmt.Println("Sum: ", sum)

	fmt.Println(minMax(23, 60))

}

//Named return
func minMax(x, y int) (max, min int) {
	if x > y{
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return

}




