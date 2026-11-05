package result

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"

func NewOversize(r *client.OversizeReport) *Oversize {
	s := &OversizeSummary{Files: len(r.Files)}

	for _, f := range r.Files {
		s.Chunks += len(f.Chunks)

		for _, c := range f.Chunks {
			if c.EmbedTokens > r.EmbedAllowance {
				s.OverEmbedder++
			}
		}
	}

	return &Oversize{Report: r, Summary: s}
}
