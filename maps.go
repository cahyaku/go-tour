package main

import "fmt"

func createLocationMap() {
	m := make(map[string]Vertex)

	m["Bell Labs"] = Vertex{
		0, 0, 40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

func printLocationMap() {
	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			0, 0, 40.68433, -74.39967,
		},
		"Google": Vertex{
			0, 0, 37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}
