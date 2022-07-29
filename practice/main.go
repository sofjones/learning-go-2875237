package main

import (
	"fmt"
)


func main() {
	colors := []string{"Burgundy", "Navy", "Majenta"}
	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}
	
	value := 0
	for value < 2 {
		fmt.Println(value)
		value++
	}
	
	for i := range colors {
		if colors[i]  == "Navy" {
			goto theEnd
		}
		fmt.Println(colors[i])
	}

	for _, color := range colors {
		fmt.Println(color)
	}
	
theEnd:
	fmt.Println("end of program")
}
