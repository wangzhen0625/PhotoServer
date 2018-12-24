package main

import (
	"fmt"
	"time"
)

func TimeCost(str string, start time.Time) {
	terminal := time.Since(start)
	fmt.Println(str, terminal)
}
