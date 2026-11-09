package test_article

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/segment"
)

func innerArticle(name string) string {
	parts := segment.Segments(name)
	last := len(parts) - 1

	for i, part := range parts {
		if i > 0 && i < last && constant.Articles[part] {
			return part
		}
	}

	return ""
}
