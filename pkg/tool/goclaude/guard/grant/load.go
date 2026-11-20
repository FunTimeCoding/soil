package grant

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
)

func (g *Grant) load() {
	if g.loaded || g.session == "" {
		return
	}

	g.loaded = true
	x, cancel := context.WithTimeout(
		context.Background(),
		constant.GrantLookupTimeout,
	)
	defer cancel()
	r, e := g.client.GetSessionLabel(x, g.session)

	if e != nil {
		g.failure = e

		return
	}

	for _, l := range r.Labels {
		if l.Key == constant.GitWriteLabel {
			g.repositories = append(g.repositories, resolved(l.Value))
		}
	}
}
