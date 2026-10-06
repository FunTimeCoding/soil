package rerank

import (
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
)

func NewEnvironment() (*Reranker, error) {
	return NewModel(
		ModelByName(environment.Required(constant.RerankModelEnvironment)),
	)
}
