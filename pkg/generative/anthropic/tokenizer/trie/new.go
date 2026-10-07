package trie

import "github.com/funtimecoding/soil/pkg/generative/types/trie_node"

func New() *Trie {
	return &Trie{root: trie_node.New()}
}
