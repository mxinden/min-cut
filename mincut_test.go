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
