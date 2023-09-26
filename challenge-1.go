package main

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

// Challenge 1.1: Represent the graph
// Challenge 1.2: Use a search method to figure out if there exists a route between PHX and BKK

// Questions to consider:
// -- What is the best way to represent this graph?
// -- How do we implement the above?
// -- What searching algorithm should we use?
// -- Once done, was this the most efficient way to search for this route?
