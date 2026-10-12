package gosecret

import (
	"maps"
	"strings"
)

func decodedMatches(
	path string,
	values map[string]string,
) (bool, error) {
	decoded, e := ReadDecoded(path)

	if e != nil {
		return false, e
	}

	if decoded == nil {
		return true, nil
	}

	trimmed := map[string]string{}

	for k, v := range values {
		trimmed[k] = strings.TrimRight(v, "\n")
	}

	return maps.Equal(trimmed, decoded), nil
}
