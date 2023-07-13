package Graphs

import "errors"

type vertex struct {
	key        string
	neighbours map[string]*neighbour
}

type neighbour struct {
	weight float64
	vertex *vertex
}

type Graph struct {
	vertices map[string]*vertex
}

// initialize a new graph
func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string]*vertex),
	}
}

// initializes a new node and sets its key
func NewVertex(key string) *vertex {
	return &vertex{
		key:        key,
		neighbours: make(map[string]*neighbour),
	}
}

// Adds a new vertex to the graph
func (g *Graph) AddVertex(key string) (*vertex, error) {
	if _, ok := g.vertices[key]; !ok {
		g.vertices[key] = NewVertex(key)
		return g.GetVertex(key), nil
	}
	return g.GetVertex(key), errors.New("vertex already exists")
}

// Returns a vertex from the graph
func (g *Graph) GetVertex(key string) *vertex {
	if _, ok := g.vertices[key]; !ok {
		return nil
	}

	return g.vertices[key]
}

// Removes a vertex from the graph
func (g *Graph) RemoveVertex(key string) (bool, error) {
	if _, ok := g.vertices[key]; !ok {
		return false, errors.New("vertex does not exist")
	}

	delete(g.vertices, key)

	for _, vertex := range g.vertices {
		delete(vertex.neighbours, key)
	}

	return true, nil
}

// Adds an edge between two vertices
func (g *Graph) AddEdge(src, dst string, weight float64) (bool, error) {
	from := g.GetVertex(src)
	if from == nil {
		return false, errors.New("source vertex does not exist")
	}

	to := g.GetVertex(dst)
	if to == nil {
		return false, errors.New("destination vertex does not exist")
	}

	from.AddNeighbour(to, weight)

	return true, nil
}

// Removes an edge between two nodes
func (g *Graph) RemoveEdge(src, dst string) (bool, error) {
	from := g.GetVertex(src)
	to := g.GetVertex(dst)

	if from == nil {
		return false, errors.New("source vertex does not exist")
	}

	if to == nil {
		return false, errors.New("destination vertex does not exist")
	}

	delete(from.neighbours, dst)
	delete(to.neighbours, src)
	return true, nil
}

// Add a new neighbour to a vertex
func (v *vertex) AddNeighbour(vertex *vertex, weight float64) {
	v.neighbours[vertex.key] = &neighbour{
		weight: weight,
		vertex: vertex,
	}

	vertex.neighbours[v.key] = &neighbour{
		weight: weight,
		vertex: v,
	}
}
