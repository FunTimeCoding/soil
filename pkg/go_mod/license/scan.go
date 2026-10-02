package license

import (
	"github.com/google/licensecheck"
	"slices"
)

func Scan(text string) []string {
	var result []string

	for _, m := range licensecheck.Scan([]byte(text)).Match {
		if !slices.Contains(result, m.ID) {
			result = append(result, m.ID)
		}
	}

	return result
}
