package main

import (
	"fmt"
	"strconv"
)

func main() {
	_, err := strconv.Atoi("1")
	fmt.Print(err)
}
