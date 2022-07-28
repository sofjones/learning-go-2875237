package main

import (
	"fmt"
)

func main() {
	var dogs [4]string
	dogs[0] = "labradoodle"
	dogs[1] = "cavapoo"
	dogs[2] = "golden retriever"

	var places = []string{"san diego", "newport beach", "victoria"}
	//fmt.Println(places[1])
	places = append(places, "petaluma")

	//fmt.Println("places been: ", len(places))

	provinces := make(map[string]string)
	provinces["BC"] = "British Columbia"
	provinces["NB"] = "New Brunswick"
	provinces["AB"] = "Alberta"

}

// Thank you all for being so welcoming! My name is Sofia, I'm new to the Secret Scanning Experiences team. I live in Southern California and love hiking, soccer, cycling, puzzles and learning about how things work. I’m wildly happy to be here and am looking forward to meeting everyone.