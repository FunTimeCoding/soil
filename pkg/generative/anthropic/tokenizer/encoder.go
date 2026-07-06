package tokenizer

import "github.com/funtimecoding/soil/pkg/generative/anthropic/tokenizer/trie"

type Encoder struct {
	trie *trie.Trie
}
