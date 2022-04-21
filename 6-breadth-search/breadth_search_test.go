package breadth_search_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isSeller(name string) bool {
	return name == "thom"
}

func BreadthSearch(graph map[string][]string, name string) string {
	queue := []string{name}
	visited := make(map[string]bool)

loop:
	for {
		if len(queue) == 0 {
			break loop
		}

		person := queue[0]
		queue = queue[1:]

		if !visited[person] {
			if isSeller(person) {
				return person
			}

			visited[person] = true
			queue = append(queue, graph[person]...)
		}
	}

	return ""
}

func TestBreadhSearch(t *testing.T) {
	graph := make(map[string][]string)

	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	assert.Equal(t, BreadthSearch(graph, "you"), "thom", "thom should be seller")
	assert.Equal(t, BreadthSearch(graph, "thom"), "thom", "thom should be seller")
}
