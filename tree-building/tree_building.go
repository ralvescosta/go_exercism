package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(b []Record) (*Node, error) {
	sort.Slice(b, func(i, j int) bool {
		return b[i].ID < b[j].ID
	})

	var nodes = map[int]*Node{}
	for i, r := range b {
		if r.ID != i {
			return nil, fmt.Errorf("missing record")
		}

		if r.Parent > 0 && r.ID <= r.Parent {
			return nil, fmt.Errorf("parent ID cannot be greater than record")
		}

		nodes[r.ID] = &Node{ID: r.ID}
		if r.ID != 0 {
			nodes[r.Parent].Children = append(nodes[r.Parent].Children, nodes[r.ID])
		}
	}
	return nodes[0], nil
}
