package service

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/types/chunk_preview"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/types/preview_chunk"
	"strings"
)

func (s *Service) Preview(
	path string,
	body string,
) *chunk_preview.Preview {
	result := chunk_preview.New(
		s.reranker.Name(),
		s.reranker.SequenceLength(),
		s.reranker.Allowance(),
		s.counter.Allowance(),
	)
	source := strings.HasSuffix(path, soil.GoExtension)

	for i, c := range chunk.Document(body, path, s.reranker) {
		first, last := lineRange(body, c)
		piece := c.Length != len(c.Text)
		p := preview_chunk.New(
			i,
			first,
			last,
			len(c.Text),
			s.reranker.Count(c.Text),
			s.counter.Count(c.Text),
			piece,
		)

		if !source && !piece && p.Tokens > result.Allowance {
			p.SetCut(s.headingFix(path, body, c))
		}

		result.Chunks = append(result.Chunks, p)
	}

	if !source {
		result.Sections = s.previewSections(body)
	}

	return result
}
