package example

import "sort"

func DescendantShadow() {
	n := 0
	items := []int{3, 1, 2}
	sort.Slice(
		items,
		func(i, j int) bool {
			n++

			return items[i] > items[j]
		},
	)
	_ = n
}
