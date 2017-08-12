package mincut

type graph struct {
	n, e  int
	edges []edge
}

type edge struct {
	a, b int
}
