package query

import "fmt"

func Quote(v []string) []string {
	result := make([]string, 0, len(v))

	for _, e := range v {
		result = append(result, fmt.Sprintf("'%s'", e))
	}

	return result
}
