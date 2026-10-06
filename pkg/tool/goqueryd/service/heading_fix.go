package service

import (
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"strings"
)

func (s *Service) headingFix(
	path string,
	body string,
	c chunk.Chunk,
) (int, []int) {
	end := min(len(body), c.Position+2*constant.ChunkSize)

	for level := constant.DeepestHeadingLevel; level >= 1; level-- {
		heading := join.Empty(
			strings.Repeat(constant.HeadingMark, level),
			constant.ProbeHeading,
		)
		var lines []int

		for offset := c.Position + 1; offset < c.Position+c.Length; offset++ {
			if body[offset-1] != '\n' ||
				strings.HasPrefix(body[offset:], constant.TableRowPrefix) {
				continue
			}

			first := chunk.Document(
				join.Empty(body[c.Position:offset], heading, body[offset:end]),
				path,
				s.reranker,
			)[0]

			if s.reranker.Count(first.Text) <= s.reranker.Allowance() {
				lines = append(
					lines,
					strings.Count(body[:offset], stringConstant.Unix)+1,
				)
			}
		}

		if len(lines) > 0 {
			return level, lines
		}
	}

	return 0, nil
}
