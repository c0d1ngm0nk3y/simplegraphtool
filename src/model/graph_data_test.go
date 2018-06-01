package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEdge(t *testing.T) {
	e := Edge{"A", "B", 42}
	assert.Equal(t, "A", e.source)
	assert.Equal(t, "B", e.destination)
	assert.Equal(t, 42, e.weight)
}

func TestEmptyGraph(t *testing.T) {
	g := NewGraphData()

	assert.Len(t, g.GetNodes(), 0)
}

func TestAddOneEdge(t *testing.T) {
	g := NewGraphData()
	g.Add(Edge{"A", "B", 1})

	nodes := g.GetNodes()
	assert.Len(t, nodes, 2)
	assert.Contains(t, nodes, "A")
	assert.Contains(t, nodes, "B")
}

func TestAddOneEdgeGetEdges(t *testing.T) {
	g := NewGraphData()
	g.Add(Edge{"A", "B", 1})

	edges := g.GetEdges()
	assert.Len(t, edges, 1)

	edge := edges[0]
	assert.Equal(t, edge.source, "A")
	assert.Contains(t, edge.destination, "B")
}

func TestAddTwoEdgesWithCommonSource(t *testing.T) {
	g := NewGraphData()
	g.Add(Edge{"A", "B", 1})
	g.Add(Edge{"A", "C", 3})

	nodes := g.GetNodes()
	assert.Len(t, nodes, 3)
	assert.Contains(t, nodes, "C")
}

func TestGraphAsStringInDotFormatNodes(t *testing.T) {
	g := NewGraphData()
	g.Add(Edge{"Node A", "Node B", 1})
	g.Add(Edge{"Node A", "Node C", 3})

	assert.Contains(t, g.String(), "Node A;")
	assert.Contains(t, g.String(), "Node C;")
}

func TestGraphAsStringInDotFormatEdges(t *testing.T) {
	g := NewGraphData()
	g.Add(Edge{"Node A", "Node B", 1})
	g.Add(Edge{"Node A", "Node C", 3})

	assert.Contains(t, g.String(), "Node A->Node C;")
}

func TestParseEdgeWithEmptyString(t *testing.T) {
	success, _ := ParseEdge("")

	assert.False(t, success)
}

func TestParseEdgeWithSomeString(t *testing.T) {
	success, _ := ParseEdge("soomom")

	assert.False(t, success)
}

func TestParseEdgeWithError(t *testing.T) {
	success, _ := ParseEdge("A,B,x")

	assert.False(t, success)
}

func TestParseEdgeWithMinimal(t *testing.T) {
	success, e := ParseEdge("A,B,1")

	assert.True(t, success)
	assert.Equal(t, "A", e.source)
	assert.Equal(t, "B", e.destination)
	assert.Equal(t, 1, e.weight)
}
