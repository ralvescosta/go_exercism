package main

import (
	"fmt"
	"gigasecond"
	"time"
)

func main() {
	t, _ := time.Parse("2006-01-02", "2011-04-25")
	fmt.Println(t)
	fmt.Println(gigasecond.AddGigasecond(t))
}
