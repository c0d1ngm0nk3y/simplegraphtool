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
