package model_context

import "strconv"

func numbers(identifiers []uint) []string {
	var result []string

	for _, i := range identifiers {
		result = append(result, strconv.FormatUint(uint64(i), 10))
	}

	return result
}
