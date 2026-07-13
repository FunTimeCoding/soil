package rerank

import (
	"github.com/amikos-tech/pure-tokenizers"
	"sync"
)

type Reranker struct {
	sequenceLength int
	tokenizer      *tokenizers.Tokenizer
	session        *rerankSession
	mutex          sync.Mutex
}
