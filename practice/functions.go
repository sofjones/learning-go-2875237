package main

import (
	"fmt"
)

func main() {
	total := sum(5, 10)
	fmt.Println(total)
}

func sum(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}
