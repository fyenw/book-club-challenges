package main

import (
	"fmt"
	"mark/challenge-1/algorithm"
)

func main() {
	// A list of airports
	var airports = []string{"PHX", "BKK", "OKC", "JFK", "LAX", "MEX", "EZE", "HEL", "LOS", "LAP", "LIM"}

	// A list of routes between airports
	var routes = [][2]string{
		{"PHX", "LAX"},
		{"PHX", "JFK"},
		{"JFK", "OKC"},
		{"JFK", "HEL"},
		{"JFK", "LOS"},
		{"MEX", "LAX"},
		{"MEX", "BKK"},
		{"MEX", "LIM"},
		{"MEX", "EZE"},
		{"LIM", "BKK"},
	}

	graph := algorithm.RepresentGraph(airports, routes)
	fmt.Println(graph)

	routeExists := algorithm.RouteExists("PHX", "BKK", graph)
	fmt.Println(routeExists)
}
