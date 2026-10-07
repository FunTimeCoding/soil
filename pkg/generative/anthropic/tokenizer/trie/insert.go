package trie

import "github.com/funtimecoding/soil/pkg/generative/types/trie_node"

func (t *Trie) Insert(token string) {
	n := t.root

	for i := 0; i < len(token); i++ {
		b := token[i]
		child, found := n.Children[b]

		if !found {
			child = trie_node.New()
			n.Children[b] = child
		}

		n = child
	}

	n.Terminal = true
}
