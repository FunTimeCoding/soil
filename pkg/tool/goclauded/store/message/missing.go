package message

func Missing(
	identifiers []uint,
	found []*Message,
) []uint {
	present := map[uint]bool{}

	for _, m := range found {
		present[m.Identifier] = true
	}

	var result []uint

	for _, i := range identifiers {
		if !present[i] {
			result = append(result, i)
		}
	}

	return result
}
