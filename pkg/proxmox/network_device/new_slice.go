package network_device

import "slices"

func NewSlice(v map[string]string) []*Device {
	var names []string

	for k := range v {
		names = append(names, k)
	}

	slices.SortFunc(
		names,
		func(
			a string,
			b string,
		) int {
			return suffix(a) - suffix(b)
		},
	)
	var result []*Device

	for _, n := range names {
		result = append(result, New(n, v[n]))
	}

	return result
}
