package orbit

import "fmt"

// MininumTransfers returns the minimum number of orbital transfers
// required between the objects from and to orbit.
func MininumTransfers(m *Map, from string, to string) (int, error) {
	fromMap := findNode(m, from)
	toMap := findNode(m, to)
	if fromMap == nil {
		return 0, fmt.Errorf("object %q not found in orbit map", from)
	} else if toMap == nil {
		return 0, fmt.Errorf("object %q not found in orbit map", to)
	}

	nodes := []*Map{fromMap.Parent, toMap.Parent}
	visited := make(map[string]int)
	transfers := 0
	for nodes[0] != nil || nodes[1] != nil {
		for i, node := range nodes {
			if node != nil {
				if s, prs := visited[node.Object]; prs {
					return s + transfers, nil
				}

				visited[node.Object] = transfers
				nodes[i] = node.Parent
			}
		}

		transfers++
	}

	return 0, fmt.Errorf("no transfer found")
}

func findNode(m *Map, object string) *Map {
	if m.Object == object {
		return m
	} else if m.Satellites == nil {
		return nil
	}

	for _, sat := range m.Satellites {
		if found := findNode(sat, object); found != nil {
			return found
		}
	}

	return nil
}
