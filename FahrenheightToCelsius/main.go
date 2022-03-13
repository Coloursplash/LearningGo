package main

import (
	"fmt"
	"strconv"
)

func main() {
	var fahrenheight string

	fmt.Println("Enter a temperature:")
	fmt.Scanln(&fahrenheight)

	num, _ := strconv.Atoi(fahrenheight)

	celsius := float32(num-32) / 1.8
	fmt.Println(fahrenheight, "F in celsius is", celsius, "C.")
}
