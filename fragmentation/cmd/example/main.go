package main

import (
	"fmt"

	f "developers-challenge/fragmentation"
)

func main() {
	fragments := make(map[int]f.Fragment)
	fragments[3] = f.Fragment{Data: "!", Hash: f.SimpleHash("!")}
	fragments[2] = f.Fragment{Data: "World", Hash: f.SimpleHash("World")}
	fragments[1] = f.Fragment{Data: "Hello", Hash: f.SimpleHash("Hello")}

	data, err := f.ReconstructData(fragments)
	if err != nil {
		return
	}

	fmt.Printf("Reconstructed data: %v\n", data)
}
