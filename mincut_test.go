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

	expectedGraph := graph{2, 1, []edge{{0, 1}}}

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
	{
		graph{2, 1, []edge{{0, 1}}},
		1,
	},
	{
		graph{4, 5, []edge{
			{0, 1},
			{0, 2},
			{0, 3},
			{1, 3},
			{2, 3},
		}},
		2,
	},
}

func TestMinCut(t *testing.T) {
	for _, test := range minCutTests {
		cut := repeatMinCut(test.g, 1000)
		if cut != test.cut {
			t.Fatal(fmt.Sprintf("Graph: %v expected %v, got %v", test.g, test.cut, cut))
		}
	}
}

func repeatMinCut(g graph, t int) int {
	min := minCut(g)
	for i := 1; i < t; i = i + 1 {
		newMin := minCut(g)
		if min > newMin {
			min = newMin
		}
	}
	return min
}

var getParentTests = []struct {
	subsets  []subset
	i        int
	expected int
}{
	{
		[]subset{{0, 0}},
		0,
		0,
	},
}

func TestGetParent(t *testing.T) {
	for _, test := range getParentTests {
		out := getParent(test.subsets, test.i)
		if out != test.expected {
			t.Fatal(fmt.Sprintf("Expected %v, got %v", test.expected, out))
		}
	}
}

var fileToMinCutTests = []struct {
	filePath string
	minCut   int
}{
	{"./examples/8", 2},
	{"./examples/8_2", 2},
	{"./examples/8_3", 1},
	{"./examples/40", 3},
	{"./examples/200", 17},
}

func TestFileToMinCut(t *testing.T) {
	for _, test := range fileToMinCutTests {
		graph, err := fileToGraph(test.filePath)
		if err != nil {
			t.Fatal(err)
		}
		out := repeatMinCut(graph, 100000)
		if out != test.minCut {
			t.Fatal(fmt.Sprintf("file: %v exptected %v, got %v", test.filePath, test.minCut, out))
		}
	}
}
