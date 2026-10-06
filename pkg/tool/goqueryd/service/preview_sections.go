package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service/preview_block"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service/preview_section"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/section/parser"
)

func (s *Service) previewSections(body string) []*preview_section.Section {
	var result []*preview_section.Section

	for _, x := range parser.Parse(body) {
		p := preview_section.New(
			x.Level,
			x.Title,
			x.FirstLine,
			x.LastLine,
			s.reranker.Count(x.Text),
		)

		for _, b := range x.Blocks {
			p.Blocks = append(
				p.Blocks,
				preview_block.New(
					b.Kind,
					b.FirstLine,
					b.LastLine,
					s.reranker.Count(b.Text),
				),
			)
		}

		result = append(result, p)
	}

	return result
}
