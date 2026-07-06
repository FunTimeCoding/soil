package tokenizer

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/tokenizer/trie"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/tokenizer/vocabulary"
)

func New() (*Encoder, error) {
	v, e := vocabulary.Load()

	if e != nil {
		return nil, e
	}

	t := trie.New()

	for _, token := range v.Verified {
		t.Insert(token)
	}

	return &Encoder{trie: t}, nil
}
