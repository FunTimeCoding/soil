package rerank

import (
	"encoding/json"
	"github.com/amikos-tech/pure-tokenizers"
	"os"
)

func newCounter(path string) (*tokenizers.Tokenizer, error) {
	b, e := os.ReadFile(path)

	if e != nil {
		return nil, e
	}

	var c map[string]any

	if f := json.Unmarshal(b, &c); f != nil {
		return nil, f
	}

	c["truncation"] = nil
	c["padding"] = nil
	b, e = json.Marshal(c)

	if e != nil {
		return nil, e
	}

	return tokenizers.FromBytes(b)
}
