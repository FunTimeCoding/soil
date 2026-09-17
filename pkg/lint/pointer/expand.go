package pointer

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func Expand(c *Candidate) []*Candidate {
	start := strings.Index(c.Span, "{")

	if start == -1 {
		return []*Candidate{c}
	}

	stop := strings.Index(c.Span[start:], "}")

	if stop == -1 {
		return []*Candidate{c}
	}

	inner := c.Span[start+1 : start+stop]

	if !strings.Contains(inner, constant.Comma) {
		return []*Candidate{c}
	}

	var result []*Candidate

	for _, alternative := range strings.Split(inner, constant.Comma) {
		result = append(
			result,
			Expand(
				&Candidate{
					Span: join.Empty(
						c.Span[:start],
						alternative,
						c.Span[start+stop+1:],
					),
					Link: c.Link,
				},
			)...,
		)
	}

	return result
}
