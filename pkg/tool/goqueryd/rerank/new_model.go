package rerank

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/system/join"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/types/rerank_model"
)

func NewModel(m *rerank_model.Model) (*Reranker, error) {
	sequenceLength := environment.FallbackInteger(
		constant.RerankSequenceEnvironment,
		m.SequenceLength,
	)

	if sequenceLength > m.MaximumLength {
		return nil, fmt.Errorf(
			"window %d exceeds %s maximum of %d",
			sequenceLength,
			m.Name,
			m.MaximumLength,
		)
	}

	return New(
		m.Name,
		join.Join(
			environment.Required(constant.RerankDirectoryEnvironment),
			m.Name,
		),
		sequenceLength,
	)
}
