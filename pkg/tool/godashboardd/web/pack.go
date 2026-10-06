package web

import "github.com/funtimecoding/soil/pkg/tool/godashboardd/board/layout"

func Pack(
	sections []*layout.Section,
	columns int,
) [][]*layout.Section {
	result := make([][]*layout.Section, columns)
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
