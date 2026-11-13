package layer

import "slices"

func (l *Layer) stack() []*Layer {
	var result []*Layer

	for current := l; current != nil; current = current.parent {
		result = append(result, current)
	}

	slices.Reverse(result)

	return result
}
