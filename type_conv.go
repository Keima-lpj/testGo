package main

import (
	"fmt"
	"strconv"
)

func main() {

	var timestamp string

	// string到int
	intS, _ := strconv.Atoi(timestamp)

	// string到int64
	intS64, _ := strconv.ParseInt(timestamp, 10, 64)

	var iamInt int

	// int到string
	iamIntString := strconv.Itoa(iamInt)

	var iamInt64 int64

	// int64到string
	iamInt64String := strconv.FormatInt(iamInt64, 10)

	fmt.Println(intS, intS64, iamIntString, iamInt64String)
}
