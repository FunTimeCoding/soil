package store

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/generative/ollama/generate_request"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
)

func ExpandQuery(
	query string,
	o *ollama.Client,
) ([]ExpandedQuery, error) {
	response, e := o.Generate(
		generate_request.New().Model(constant.ExpandModel).Prompt(
			join.Empty(constant.ExpandPrompt, query),
		),
	)

	if e != nil {
		return nil, e
	}

	result := parseExpandedQueries(response.Text, query)

	if len(result) == 0 {
		return []ExpandedQuery{
			{Type: "lex", Query: query},
			{Type: "vec", Query: query},
		}, nil
	}

	return result, nil
}
