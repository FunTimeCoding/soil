package rerank

import "github.com/amikos-tech/pure-tokenizers"

func (r *Reranker) measureSpecials() error {
	pair, e := r.counter.EncodePair("", "", tokenizers.WithAddSpecialTokens())

	if e != nil {
		return e
	}

	single, f := r.counter.Encode("", tokenizers.WithAddSpecialTokens())

	if f != nil {
		return f
	}

	r.pairSpecials = len(pair.IDs)
	r.singleSpecials = len(single.IDs)

	return nil
}
