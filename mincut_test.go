package mincut

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFileToGraph(t *testing.T) {
	graph, err := fileToGraph("./examples/2")
	if err != nil {
		t.Fatal(err)
	}
	expectedGraph := []node{
		{1, []int{2}},
		{2, []int{1}},
	}

	if !reflect.DeepEqual(expectedGraph, graph) {
		t.Fatal(fmt.Sprintf("Expected %v, got %v", expectedGraph, graph))
	}
}
