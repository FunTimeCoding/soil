package rerank

import (
	"github.com/amikos-tech/pure-tokenizers"
	"sync"
)

type Reranker struct {
	name           string
	sequenceLength int
	tokenizer      *tokenizers.Tokenizer
	counter        *tokenizers.Tokenizer
	pairSpecials   int
	singleSpecials int
	session        *rerankSession
	mutex          sync.Mutex
	counterMutex   sync.Mutex
}
