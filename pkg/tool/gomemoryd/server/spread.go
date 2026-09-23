package server

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/token_summary"
)

func spread(s *token_summary.Spread) server.Spread {
	return server.Spread{
		Median:      s.Median,
		Ninetieth:   s.Ninetieth,
		NinetyNinth: s.NinetyNinth,
		Maximum:     s.Maximum,
		Total:       s.Total,
	}
}
