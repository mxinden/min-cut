package mincut

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFileToGraph(t *testing.T) {
	g, err := fileToGraph("./examples/2")
	if err != nil {
		t.Fatal(err)
	}

	expectedGraph := graph{2, 1, []edge{{1, 2}}}

	if !reflect.DeepEqual(expectedGraph, g) {
		t.Fatal(fmt.Sprintf("Expected %v, got %v", expectedGraph, g))
	}
}

var minCutTests = []struct {
	g   graph
	cut int
}{
	{
		graph{0, 0, []edge{}},
		0,
	},
}

func TestMinCut(t *testing.T) {
	for _, test := range minCutTests {
		cut := minCut(test.g)
		if cut != test.cut {
			t.Fatal(fmt.Sprintf("Expected %v, got %v", test.cut, cut))
		}
	}
}
