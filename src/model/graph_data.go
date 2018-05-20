package model

import (
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
}

//GraphData is the interface to use
type GraphData interface {
	Add(edge Edge)
	GetNodes() []string
	String() string
}

//NewGraphData should be used to create graphdata
func NewGraphData() GraphData {
	g := new(graphData)
	g.nodes = make(map[string]bool)
	return g
}

//Add an edge to the graph
func (graph *graphData) Add(edge Edge) {
	graph.nodes[edge.source] = true
	graph.nodes[edge.destination] = true
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

//String puts out this graph as dot format
func (graph graphData) String() string {
	dotGraph := gographviz.NewGraph()
	for _, node := range graph.GetNodes() {
		dotGraph.AddNode("G", node, nil)
	}
	//dotGraph.AddEdge("a", "b", true, nil)
	output := dotGraph.String()
	return output
}
