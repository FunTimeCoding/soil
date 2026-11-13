package layer

import "slices"

func (l *Layer) touched() []string {
	var result []string

	for current := l; current != nil; current = current.parent {
		result = append(result, current.order...)
	}

	slices.Sort(result)

	return slices.Compact(result)
}
