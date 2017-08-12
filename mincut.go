package mincut

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
	return 0
}
