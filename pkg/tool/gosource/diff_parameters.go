package gosource

import "github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"

func diffParameters(
	module string,
	from string,
	to string,
	text bool,
) *client.GetDiffParams {
	result := &client.GetDiffParams{Module: module}

	if from != "" {
		result.From = &from
	}

	if to != "" {
		result.To = &to
	}

	if text {
		result.Text = &text
	}

	return result
}
