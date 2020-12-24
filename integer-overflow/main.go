package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "2147483648"

	// 1
	// bigValue, err := strconv.Atoi(s)

	// 2
	// bigValue, err := strconv.ParseInt(s, 10, 64)

	// 3
	bigValue, err := strconv.ParseInt(s, 10, 32)

	if err != nil {
		panic(err)
	}
	if bigValue < 0 {
		return
	}
	fmt.Printf("%v is %v\n", bigValue, int32(bigValue))
}
