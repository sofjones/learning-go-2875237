package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers")
	anInt := 4
	var p = &anInt

	fmt.Println(&anInt)
	fmt.Println(*p)
	*p = anInt / 2
	fmt.Println("set pointer value to half anInt value", *p)
	fmt.Println(anInt) //anInt value is changed
	//fmt.Printf("%d",anInt)
}
