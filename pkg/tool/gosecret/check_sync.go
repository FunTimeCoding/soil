package gosecret

import (
	"fmt"
	"reflect"
	"strings"
)

func CheckSync(
	path string,
	expected map[string]string,
) (bool, error) {
	actual, e := ReadDecoded(path)

	if e != nil {
		return false, fmt.Errorf("read decoded: %w", e)
	}

	if actual == nil {
		return false, nil
	}

	normalized := make(map[string]string, len(expected))

	for k, v := range expected {
		normalized[k] = strings.TrimRight(v, "\n")
	}

	return reflect.DeepEqual(normalized, actual), nil
}
