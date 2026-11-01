package report

import "github.com/funtimecoding/soil/pkg/tool/gosourced/types/entry"

func moves(pairs []*entry.Entry) map[string]string {
	result := map[string]string{}

	for _, p := range pairs {
		if p.From.Receiver != "" {
			continue
		}

		if from, to := p.From.Reference(), p.To.Reference(); from != to {
			result[from] = to
		}
	}

	return result
}
