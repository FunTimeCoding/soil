package selection_filter

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"maps"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"slices"
)

func (f *Filter) Node() gomponents.Node {
	s := f.selection
	coverage := make([]gomponents.Node, 0, 4)

	for _, v := range [][2]string{
		{"All", ""},
		{"Supported", constant.ShowSupported},
		{"Tracked", constant.ShowTracked},
		{"Untracked", constant.ShowUntracked},
	} {
		coverage = append(
			coverage,
			f.chip(
				v[0],
				f.count.Coverage[v[1]],
				s.WithCoverage(v[1]),
				s.Show == v[1],
			),
		)
	}

	scores := make([]gomponents.Node, 0, len(f.floors))

	for _, floor := range f.floors {
		label := "All scores"

		if floor > 0 {
			label = fmt.Sprintf("%.1f or higher", floor)
		}

		scores = append(
			scores,
			f.chip(
				label,
				f.count.Score[floor],
				s.WithMinimum(floor),
				s.Minimum == floor,
			),
		)
	}

	result := []gomponents.Node{html.Nav(coverage...), html.Nav(scores...)}

	for _, key := range slices.Sorted(maps.Keys(f.count.Labels)) {
		counts := f.count.Labels[key]
		chips := []gomponents.Node{
			html.Small(gomponents.Text(key)),
			f.chip("all", counts[""], s.WithLabel(key, ""), s.Labels[key] == ""),
		}

		for _, value := range slices.Sorted(maps.Keys(counts)) {
			if value == "" {
				continue
			}

			chips = append(
				chips,
				f.chip(
					value,
					counts[value],
					s.WithLabel(key, value),
					s.Labels[key] == value,
				),
			)
		}

		result = append(result, html.Nav(chips...))
	}

	return gomponents.Group(append(result, f.search()))
}
