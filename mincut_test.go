package mincut

import (
	"testing"
)

func TestFileToGraph(t *testing.T) {
	graph, err := fileToGraph("./examples/2")
	if err != nil {
		t.Fatal(err)
	}
}
