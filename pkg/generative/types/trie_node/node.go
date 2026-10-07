package trie_node

type Node struct {
	Children map[byte]*Node
	Terminal bool
}
