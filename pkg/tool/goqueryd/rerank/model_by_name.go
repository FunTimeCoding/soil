package rerank

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/types/rerank_model"
)

func ModelByName(name string) *rerank_model.Model {
	var known []string

	for _, m := range constant.RerankModels {
		if m.Name == name {
			return m
		}

		known = append(known, m.Name)
	}

	panic(
		fmt.Sprintf(
			"unknown rerank model %q, known: %s",
			name,
			join.CommaSpace(known),
		),
	)
}
