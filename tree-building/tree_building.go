package tree

import (
	"errors"
	"sort"
)

type Node struct {
	ID       int
	Children []*Node
}

type Record struct {
	ID, Parent int
}

func Build(records []Record) (*Node, error) {
	// Sort records by ID.
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	// Create a new node map.
	nodeMap := make(map[int]*Node)
	for i, record := range records {
		// Validate conditions of failure.
		if record.ID != i || record.Parent > record.ID || record.ID > 0 && record.Parent == record.ID {
			return nil, errors.New("Invalid Record.")
		}

		nodeMap[record.ID] = &Node{ID: record.ID}
		if record.ID > 0 {
			nodeMap[record.Parent].Children = append(nodeMap[record.Parent].Children, nodeMap[record.ID])
		}
	}

	return nodeMap[0], nil
}
