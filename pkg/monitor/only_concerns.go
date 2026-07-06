package monitor

import "github.com/funtimecoding/soil/pkg/face"

func OnlyConcerns[T face.Validator](
	v []T,
	override bool,
) []T {
	if override {
		return v
	}

	var result []T

	for _, e := range v {
		if e.HasConcerns() {
			result = append(result, e)
		}
	}

	return result
}
