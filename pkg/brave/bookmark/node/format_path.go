package node

import "github.com/funtimecoding/soil/pkg/strings/join"

func (n *Node) FormatPath() string {
	var result []string

	for _, p := range n.GetPath() {
		result = append(result, p.Name)
	}

	return join.Slash(result)
}
