package rerank

import (
	"github.com/amikos-tech/pure-tokenizers"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/types/rerank_session"
	"sync"
)

type Reranker struct {
	name           string
	sequenceLength int
	tokenizer      *tokenizers.Tokenizer
	counter        *tokenizers.Tokenizer
	pairSpecials   int
	singleSpecials int
	session        *rerank_session.Session
	mutex          sync.Mutex
	counterMutex   sync.Mutex
}
