package system_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/board_result"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
	"maps"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"slices"
)

func Filter(
	s *system_selection.Selection,
	b *board_result.Result,
) gomponents.Node {
	coverage := make([]gomponents.Node, 0, 4)

	for _, v := range [][2]string{
		{"All", ""},
		{"Supported", constant.ShowSupported},
		{"Tracked", constant.ShowTracked},
		{"Untracked", constant.ShowUntracked},
	} {
		coverage = append(
			coverage,
			facetChip(
				v[0],
				b.CoverageCounts[v[1]],
				s.WithCoverage(v[1]),
				s.Show == v[1],
			),
		)
	}

	result := []gomponents.Node{
		html.Nav(coverage...),
		html.Nav(
			facetChip(
				"All scores",
				b.ScoreCounts[0],
				s.WithMinimum(0),
				s.Minimum == 0,
			),
			facetChip(
				"9.0 or higher",
				b.ScoreCounts[constant.CriticalScore],
				s.WithMinimum(constant.CriticalScore),
				s.Minimum == constant.CriticalScore,
			),
		),
	}

	for _, key := range slices.Sorted(maps.Keys(b.LabelCounts)) {
		counts := b.LabelCounts[key]
		chips := []gomponents.Node{
			html.Small(gomponents.Text(key)),
			facetChip(
				"all",
				counts[""],
				s.WithLabel(key, ""),
				s.Labels[key] == "",
			),
		}

		for _, value := range slices.Sorted(maps.Keys(counts)) {
			if value == "" {
				continue
			}

			chips = append(
				chips,
				facetChip(
					value,
					counts[value],
					s.WithLabel(key, value),
					s.Labels[key] == value,
				),
			)
		}

		result = append(result, html.Nav(chips...))
	}

	return gomponents.Group(append(result, search(s)))
}
