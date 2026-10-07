package rerank

import (
	"fmt"
	"github.com/amikos-tech/pure-tokenizers"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/types/rerank_result"
)

func (r *Reranker) Rank(
	query string,
	documents []string,
) ([]*rerank_result.Result, error) {
	if len(documents) == 0 {
		return nil, nil
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()
	queries := make([]string, len(documents))
	texts := make([]string, len(documents))

	for i := range queries {
		queries[i] = query
		texts[i] = valid(documents[i])
	}

	encodings, e := r.tokenizer.EncodePairs(
		queries,
		texts,
		tokenizers.WithAddSpecialTokens(),
		tokenizers.WithReturnAttentionMask(),
	)

	if e != nil {
		return nil, fmt.Errorf("tokenize pairs: %w", e)
	}

	results := make([]*rerank_result.Result, len(documents))
	scoreByDocument := map[string]float64{}

	for i, encoding := range encodings {
		if score, okay := scoreByDocument[documents[i]]; okay {
			results[i] = rerank_result.New(i, score)

			continue
		}

		for j := 0; j < r.sequenceLength; j++ {
			if j < len(encoding.IDs) {
				r.session.InputIDs[j] = int64(encoding.IDs[j])
			} else {
				r.session.InputIDs[j] = 0
			}

			if j < len(encoding.AttentionMask) {
				r.session.AttentionMask[j] = int64(encoding.AttentionMask[j])
			} else {
				r.session.AttentionMask[j] = 0
			}
		}

		if f := r.session.Session.Run(); f != nil {
			return nil, fmt.Errorf("rerank inference: %w", f)
		}

		score := sigmoid(float64(r.session.OutputTensor.GetData()[0]))
		scoreByDocument[documents[i]] = score
		results[i] = rerank_result.New(i, score)
	}

	return results, nil
}
