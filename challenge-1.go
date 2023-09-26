package main

import (
	"fmt"
)

var airports = []string{"PHX", "BKK", "OKC", "JFK", "LAX", "MEX", "EZE", "HEL", "LOS", "LAP", "LIM"}

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

// Define a struct to represent the graph
type Graph struct {
	AdjacencyList map[string][]string
}

// Create a new graph
func NewGraph() *Graph {
	return &Graph{
		AdjacencyList: make(map[string][]string),
	}
}

// Add a node to the graph
func (g *Graph) AddNode(airport string) {
	g.AdjacencyList[airport] = []string{}
}

// Add an undirected edge to the graph
func (g *Graph) AddEdge(origin, destination string) {
	g.AdjacencyList[origin] = append(g.AdjacencyList[origin], destination)
	g.AdjacencyList[destination] = append(g.AdjacencyList[destination], origin)
}

func main() {
	// Create the graph
	graph := NewGraph()

	// Add nodes
	for _, airport := range airports {
		graph.AddNode(airport)
	}

	// Add edges
	for _, route := range routes {
		graph.AddEdge(route[0], route[1])
	}

	// Print the adjacency list
	for airport, connections := range graph.AdjacencyList {
		fmt.Printf("%s -> %v\n", airport, connections)
	}
}
