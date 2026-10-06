package pointer

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/markup/heading"
	"github.com/funtimecoding/soil/pkg/strings/distance"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"slices"
)

func nearest(
	fragment string,
	headings []*heading.Heading,
) string {
	if len(headings) == 0 {
		return ""
	}

	ranked := slices.Clone(headings)
	slices.SortStableFunc(
		ranked,
		func(
			a *heading.Heading,
			b *heading.Heading,
		) int {
			return distance.Levenshtein(fragment, a.Slug) -
				distance.Levenshtein(fragment, b.Slug)
		},
	)
	var result []string

	for _, h := range ranked[:min(len(ranked), constant.NearestLimit)] {
		result = append(
			result,
			fmt.Sprintf(constant.NearestFormat, h.Slug, h.Text),
		)
	}

	return join.Empty(constant.NearestPrefix, join.CommaSpace(result))
}
