package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	rand.Seed(time.Now().Unix())
	dayOfWeek := rand.Intn(7) + 1
	fmt.Println("day of week: ", dayOfWeek)	
	var result string
	switch dayOfWeek {
	case 1:
		result = "It's sunday!"
	case 2:
		result = "it's monday"
	case 3:
		result = "its tuesday"
	default:
		result = "another day"
	}
	fmt.Println(result)

}
