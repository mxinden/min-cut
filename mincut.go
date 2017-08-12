package mincut

import (
	"math/rand"
)

type graph struct {
	n, e  int
	edges []edge
}

type edge struct {
	a, b int
}

type subset struct {
	parent, rank int
}

func minCut(g graph) int {
	subsets := make([]subset, g.n)
	for i := 0; i < g.n; i = i + 1 {
		subsets[i].parent = i
		subsets[i].rank = 0
	}

	amountNodes := g.n

	for amountNodes > 2 {
		e := rand.Intn(len(g.edges))

		subsetA := getParent(subsets, g.edges[e].a)
		subsetB := getParent(subsets, g.edges[e].b)

		if subsetA == subsetB {
			continue
		}

		amountNodes = amountNodes - 1
		contract(subsets, subsetA, subsetB)

	}

	edgesCut := 0
	for i := 0; i < g.e; i = i + 1 {
		subsetA := getParent(subsets, g.edges[i].a)
		subsetB := getParent(subsets, g.edges[i].b)
		if subsetA != subsetB {
			edgesCut = edgesCut + 1
		}
	}

	return edgesCut
}

func getParent(subsets []subset, i int) int {
	if subsets[i].parent != i {
		subsets[i].parent = getParent(subsets, subsets[i].parent)
	}
	return subsets[i].parent
}

func contract(subsets []subset, subsetA, subsetB int) {
	parentA := getParent(subsets, subsetA)
	parentB := getParent(subsets, subsetB)

	if subsets[parentA].rank < subsets[parentB].rank {
		subsets[parentA].parent = parentB
	} else if subsets[parentA].rank > subsets[parentB].rank {
		subsets[parentB].parent = parentA
	} else {
		subsets[parentB].parent = parentA
		subsets[parentA].rank = subsets[parentA].rank + 1
	}
}
