package mincut

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type node struct {
	id     int
	connTo []int
}

func fileToGraph(filePath string) (graph, error) {
	nodes := []node{}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return graph{}, err
	}

	stringNodes := strings.Split(string(content), "\n")
	for _, stringNode := range stringNodes {
		if stringNode == "" {
			continue
		}
		n, err := parseNode(stringNode)
		if err != nil {
			return graph{}, err
		}
		nodes = append(nodes, n)
	}

	edges := nodesToEdges(nodes)
	graph := graph{len(nodes), len(edges), edges}
	return graph, nil
}

func parseNode(s string) (node, error) {
	n := node{}
	ids := strings.Split(s, "\t")
	for i, stringId := range ids {
		if stringId == "" {
			continue
		}

		id, err := strconv.Atoi(stringId)
		if err != nil {
			return n, err
		}

		if i == 0 {
			n.id = id
			continue
		}

		n.connTo = append(n.connTo, id)
	}
	return n, nil
}

func nodesToEdges(nodes []node) []edge {
	edges := []edge{}
	for _, n := range nodes {
		for _, c := range n.connTo {
			// Already present in edges
			if n.id > c {
				continue
			}

			edges = append(edges, edge{n.id, c})
		}
	}
	return edges
}
