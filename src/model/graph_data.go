package model

import (
	"strconv"
	"strings"

	"github.com/awalterschulze/gographviz"
)

//Edge points from source to destination
type Edge struct {
	source      string
	destination string
	weight      int
}

//GraphData representing a graph
type graphData struct {
	nodes map[string]bool
	edges []Edge
}

//GraphData is the interface to use
type GraphData interface {
	Add(edge Edge)
	GetNodes() []string
	GetEdges() []Edge
	String() string
}

//NewGraphData should be used to create graphdata
func NewGraphData() GraphData {
	g := new(graphData)
	g.nodes = make(map[string]bool)
	return g
}

//ParseEdge from a csv line
func ParseEdge(s string) (bool, Edge) {
	var e Edge
	tokens := strings.Split(s, ",")

	if len(tokens) == 3 {
		weight, error := strconv.Atoi(tokens[2])
		if error == nil {
			e = Edge{tokens[0], tokens[1], weight}
			return true, e
		}
	}

	return false, e
}

//Add an edge to the graph
func (graph *graphData) Add(edge Edge) {
	graph.nodes[edge.source] = true
	graph.nodes[edge.destination] = true
	graph.edges = append(graph.edges, edge)
}

//GetNodes of this graph
func (graph graphData) GetNodes() []string {
	keys := make([]string, len(graph.nodes))

	i := 0
	for k := range graph.nodes {
		keys[i] = k
		i++
	}

	return keys
}

//GetEdges of this graph
func (graph graphData) GetEdges() []Edge {
	return graph.edges
}

//String puts out this graph as dot format
func (graph graphData) String() string {
	dotGraph := gographviz.NewGraph()
	for _, node := range graph.GetNodes() {
		dotGraph.AddNode("G", node, nil)
	}

	for _, edge := range graph.GetEdges() {
		dotGraph.AddEdge(edge.source, edge.destination, true, nil)
	}
	output := dotGraph.String()
	return output
}
