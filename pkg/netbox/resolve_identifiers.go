package netbox

func resolveIdentifiers[T any](
	names []string,
	lookup func(string) (T, error),
	identifier func(T) int32,
) ([]int32, error) {
	result := []int32{}

	for _, n := range names {
		v, e := lookup(n)

		if e != nil {
			return nil, e
		}

		result = append(result, identifier(v))
	}

	return result, nil
}
