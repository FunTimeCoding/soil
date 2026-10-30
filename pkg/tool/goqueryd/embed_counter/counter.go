package embed_counter

import (
	"github.com/amikos-tech/pure-tokenizers"
	"sync"
)

type Counter struct {
	tokenizer *tokenizers.Tokenizer
	specials  int
	mutex     sync.Mutex
}
