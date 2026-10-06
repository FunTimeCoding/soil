package mock_reranker

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (r *Reranker) Name() string {
	return constant.MockRerankerName
}
