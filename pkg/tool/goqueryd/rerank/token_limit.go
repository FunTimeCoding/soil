package rerank

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (r *Reranker) TokenLimit() int {
	return r.sequenceLength - r.pairSpecials - constant.RerankQueryReserve
}
