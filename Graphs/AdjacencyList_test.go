package Graphs

import (
	"testing"
)

func TestGraph(t *testing.T) {
	t.Parallel()

	t.Run("Create a graph", func(t *testing.T) {
		t.Parallel()

		g := NewGraph()

		if g == nil {
			t.Error("Graph not created")
		}
	})

	t.Run("Create a vertex", func(t *testing.T) {
		t.Parallel()

		v := NewVertex("a")

		if v.key != "a" {
			t.Errorf("Expected key to be a, got %s", v.key)
		}
	})

	t.Run("Add a vertex", func(t *testing.T) {
		t.Parallel()

		g := NewGraph()

		v, err := g.AddVertex("a")

		if err != nil {
			t.Error("Vertex not added")
		}

		if v.key != "a" {
			t.Errorf("Expected key to be a, got %s", v.key)
		}

		if len(g.vertices) != 1 {
			t.Errorf("Expected length of vertices to be 1, got %d", len(g.vertices))
		}

		if g.vertices["a"] == nil {
			t.Errorf("Expected vertex with key a to exist, got %v", g.vertices["a"])
		}
	})

	t.Run("Get a vertex", func(t *testing.T) {
		t.Parallel()

		g := NewGraph()

		g.AddVertex("a")

		v := g.GetVertex("a")

		if v == nil {
			t.Errorf("Expected vertex to not be nil, got %v", v)
		}

		if v.key != "a" {
			t.Errorf("Expected key to be a, got %s", v.key)
		}
	})

	t.Run("Remove a vertex", func(t *testing.T) {
		t.Parallel()

		g := NewGraph()

		g.AddVertex("a")

		ok, err := g.RemoveVertex("a")

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if !ok {
			t.Errorf("Expected vertex with key a to be removed, got %v", ok)
		}

		if len(g.vertices) != 0 {
			t.Errorf("Expected length of vertices to be 0, got %d", len(g.vertices))
		}

		if g.vertices["a"] != nil {
			t.Errorf("Expected vertex with key a to not exist, got %v", g.vertices["a"])
		}
	})

	t.Run("Add an edge", func(t *testing.T) {
		t.Parallel()

		g := NewGraph()

		g.AddVertex("a")
		g.AddVertex("b")

		_, err := g.AddEdge("a", "b", 1)

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(g.vertices["a"].neighbours) != 1 {
			t.Errorf("Expected a to have 1 neighbour, got %d", len(g.vertices["a"].neighbours))
		}

		if len(g.vertices["b"].neighbours) != 1 {
			t.Errorf("Expected b to have 1 neighbour, got %d", len(g.vertices["b"].neighbours))
		}
	})

	t.Run("Remove an edge", func(t *testing.T) {
		t.Parallel()

		g := NewGraph()

		g.AddVertex("a")
		g.AddVertex("b")

		g.AddEdge("a", "b", 1)
		_, err := g.RemoveEdge("a", "b")

		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		if len(g.vertices["a"].neighbours) != 0 {
			t.Errorf("Expected a to have 0 neighbours, got %d", len(g.vertices["a"].neighbours))
		}

		if len(g.vertices["b"].neighbours) != 0 {
			t.Errorf("Expected b to have 0 neighbours, got %d", len(g.vertices["b"].neighbours))
		}
	})

	t.Run("Add a neighbour", func(t *testing.T) {
		t.Parallel()

		g := NewGraph()

		a, _ := g.AddVertex("a")
		b, _ := g.AddVertex("b")

		a.AddNeighbour(b, 1)

		if len(a.neighbours) != 1 {
			t.Errorf("Expected a to have 1 neighbour, got %d", len(a.neighbours))
		}

		if len(b.neighbours) != 1 {
			t.Errorf("Expected b to have 1 neighbour, got %d", len(b.neighbours))
		}

		if a.neighbours[b.key].vertex.key != "b" {
			t.Errorf("Expected a's neighbour to be b, got %s", a.neighbours[b.key].vertex.key)
		}

		if b.neighbours[a.key].vertex.key != "a" {
			t.Errorf("Expected b's neighbour to be a, got %s", b.neighbours[a.key].vertex.key)
		}
	})
}
