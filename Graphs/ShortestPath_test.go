package Graphs

import "testing"

type edge struct {
	src    string
	dest   string
	weight int
}

type TestCase struct {
	vertices []string
	edges    []edge
	source   string
	expected map[string]float64
}

func TestShortestPath(t *testing.T) {

	t.Run("Shortest path from source to all other vertices", func(t *testing.T) {
		t.Parallel()

		testCase := TestCase{
			vertices: []string{
				"kampala", "jinja", "kisumu", "kitale", "eldoret", "nakuru", "nyeri", "nairobi", "mombasa", "arusha", "tanga", "dodoma", "morogoro", "dar-es-salam", "lindi",
			},
			edges: []edge{
				{"kampala", "jinja", 80},
				{"jinja", "kisumu", 230},
				{"kisumu", "kitale", 169},
				{"kisumu", "nakuru", 185},
				{"kisumu", "nairobi", 346},
				{"kitale", "eldoret", 71},
				{"eldoret", "nakuru", 156},
				{"nakuru", "nyeri", 163},
				{"nakuru", "nairobi", 157},
				{"nairobi", "nyeri", 150},
				{"nairobi", "mombasa", 488},
				{"nairobi", "arusha", 269},
				{"mombasa", "arusha", 389},
				{"mombasa", "tanga", 203},
				{"arusha", "tanga", 436},
				{"arusha", "dodoma", 413},
				{"dodoma", "morogoro", 264},
				{"tanga", "morogoro", 333},
				{"tanga", "dar-es-salam", 356},
				{"dar-es-salam", "morogoro", 194},
				{"dar-es-salam", "lindi", 457},
			},
			source: "nairobi",
			expected: map[string]float64{
				// direct distances
				"nairobi": 0,
				"nyeri":   150,
				"nakuru":  157,
				"kisumu":  342,
				"mombasa": 488,
				"arusha":  269,
				// indirect distances
				// ke
				"eldoret": 313,
				"kitale":  384,
				// ug
				"jinja":   572,
				"kampala": 652,
				// tz
				"tanga":        691,
				"dodoma":       682,
				"morogoro":     946,
				"dar-es-salam": 1047,
				"lindi":        1504,
			},
		}

		g := NewGraph()

		for _, vertex := range testCase.vertices {
			g.AddVertex(vertex)
		}

		for _, edge := range testCase.edges {
			g.AddEdge(edge.src, edge.dest, float64(edge.weight))
		}

		distances := g.ShortestPath(g.vertices[testCase.source])

		for key, value := range distances {
			if value != testCase.expected[key] {
				t.Errorf("Expected %s - %s to be %f, got %f", testCase.source, key, testCase.expected[key], value)
			}
		}
	})
}
