package gosecret

import (
	"fmt"
	"os"
	"strings"
)

func ReadDecoded(path string) (map[string]string, error) {
	b, e := os.ReadFile(path)

	if e != nil {
		if os.IsNotExist(e) {
			return nil, nil
		}

		return nil, fmt.Errorf("read file: %w", e)
	}

	result := make(map[string]string)
	var currentKey string
	var valueLines []string
	save := func() {
		if currentKey != "" {
			result[currentKey] = strings.TrimRight(
				strings.Join(valueLines, "\n"),
				"\n",
			)
		}
	}

	for _, line := range strings.Split(string(b), "\n") {
		if strings.HasPrefix(line, "=== ") && strings.HasSuffix(
			line,
			" ===",
		) {
			save()
			currentKey = line[4 : len(line)-4]
			valueLines = nil
		} else if currentKey != "" {
			valueLines = append(valueLines, line)
		}
	}

	save()

	return result, nil
}
