package mincut

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func fileToGraph(filePath string) ([]node, error) {
	graph := []node{}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return graph, err
	}

	stringNodes := strings.Split(string(content), "\n")
	for _, stringNode := range stringNodes {
		if stringNode == "" {
			continue
		}
		n, err := parseNode(stringNode)
		if err != nil {
			return graph, err
		}
		graph = append(graph, n)
	}
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
