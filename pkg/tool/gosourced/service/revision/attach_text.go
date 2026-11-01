package revision

import "github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"

func attachText(
	symbols []*symbol.Symbol,
	files map[string][]byte,
) {
	for _, s := range symbols {
		b := files[s.File]

		if s.Start < 0 || s.End > len(b) || s.Start > s.End {
			continue
		}

		s.Text = string(b[s.Start:s.End])
	}
}
