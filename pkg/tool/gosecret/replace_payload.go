package gosecret

import (
	"encoding/base64"
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"sort"
	"strings"
)

func ReplacePayload(
	m []byte,
	decoded map[string]string,
) ([]byte, error) {
	lines := strings.Split(string(m), "\n")
	var result []string
	found := false
	i := 0

	for ; i < len(lines); i++ {
		result = append(result, lines[i])

		if lines[i] == "data:" {
			found = true
			i++

			break
		}
	}

	if !found {
		return nil, fmt.Errorf("no data block in manifest")
	}

	for ; i < len(lines); i++ {
		if !strings.HasPrefix(lines[i], " ") {
			break
		}
	}

	keys := make([]string, 0, len(decoded))

	for k := range decoded {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		v := decoded[k]

		if strings.Contains(v, "\n") {
			v = join.Empty(v, "\n")
		}

		result = append(
			result,
			fmt.Sprintf(
				"  %s: %s",
				k,
				base64.StdEncoding.EncodeToString([]byte(v)),
			),
		)
	}

	result = append(result, lines[i:]...)

	return []byte(strings.Join(result, "\n")), nil
}
