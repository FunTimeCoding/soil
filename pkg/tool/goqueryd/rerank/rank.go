package rerank

import (
	"fmt"
	"github.com/amikos-tech/pure-tokenizers"
)

func (r *Reranker) Rank(
	query string,
	documents []string,
) ([]Result, error) {
	if len(documents) == 0 {
		return nil, nil
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()
	queries := make([]string, len(documents))

	for i := range queries {
		queries[i] = query
	}

	encodings, e := r.tokenizer.EncodePairs(
		queries,
		documents,
		tokenizers.WithAddSpecialTokens(),
		tokenizers.WithReturnAttentionMask(),
	)

	if e != nil {
		return nil, fmt.Errorf("tokenize pairs: %w", e)
	}

	results := make([]Result, len(documents))
	scoreByDocument := map[string]float64{}

	for i, encoding := range encodings {
		if score, okay := scoreByDocument[documents[i]]; okay {
			results[i] = Result{Index: i, Score: score}

			continue
		}

		for j := 0; j < r.sequenceLength; j++ {
			if j < len(encoding.IDs) {
				r.session.inputIDs[j] = int64(encoding.IDs[j])
			} else {
				r.session.inputIDs[j] = 0
			}

			if j < len(encoding.AttentionMask) {
				r.session.attentionMask[j] = int64(encoding.AttentionMask[j])
			} else {
				r.session.attentionMask[j] = 0
			}
		}

		if f := r.session.session.Run(); f != nil {
			return nil, fmt.Errorf("rerank inference: %w", f)
		}

		score := sigmoid(float64(r.session.outputTensor.GetData()[0]))
		scoreByDocument[documents[i]] = score
		results[i] = Result{Index: i, Score: score}
	}

	return results, nil
}
