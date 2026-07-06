package web

import "github.com/funtimecoding/soil/pkg/tool/godashboardd/board"

func pack(
	sections []*board.Section,
	columns int,
) [][]*board.Section {
	result := make([][]*board.Section, columns)
	weights := make([]int, columns)

	for _, v := range sections {
		shortest := 0

		for i, w := range weights {
			if w < weights[shortest] {
				shortest = i
			}
		}

		result[shortest] = append(result[shortest], v)
		weights[shortest] += weight(v)
	}

	return result
}
