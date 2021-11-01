package tree

import (
	"errors"
	"sort"
)

// Define the Record type
type Record struct {
	ID     int
	Parent int
}

// Define the Node type
type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	n := make(map[int]*Node)

	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	for i, r := range records {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
			return nil, errors.New("record has bad parent or out of sequence")
		}

		n[r.ID] = &Node{ID: r.ID}

		if r.ID != 0 {
			n[r.Parent].Children = append(n[r.Parent].Children, n[r.ID])
		}
	}

	return n[0], nil

}
