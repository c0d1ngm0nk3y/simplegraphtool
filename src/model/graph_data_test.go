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

func TestAddTwoEdgesWithCommonSource(t *testing.T) {
	g := NewGraphData()
	g.Add(Edge{"A", "B", 1})
	g.Add(Edge{"A", "C", 3})

	nodes := g.GetNodes()
	assert.Len(t, nodes, 3)
	assert.Contains(t, nodes, "C")
}
