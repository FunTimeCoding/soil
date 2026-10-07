package trie_node

func New() *Node {
	return &Node{Children: map[byte]*Node{}}
}
