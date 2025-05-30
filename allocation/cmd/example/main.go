package main

import (
	"developers-challenge/allocation"
	"fmt"
)

func main() {
	fragments := 5
	risks := []int{10, 20, 30}
	minimized := allocation.DistributeFragments(risks, fragments)

	fmt.Printf(
		"minimized risk: %v (%v fragments, distributed across: %v data centers\n",
		minimized,
		fragments,
		risks,
	)
}
