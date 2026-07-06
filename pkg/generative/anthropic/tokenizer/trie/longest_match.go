package trie

func (t *Trie) LongestMatch(
	text string,
	position int,
) int {
	node := t.root
	best := 0

	for i := position; i < len(text); i++ {
		child, found := node.children[text[i]]

		if !found {
			break
		}

		node = child

		if node.terminal {
			best = i - position + 1
		}
	}

	return best
}
