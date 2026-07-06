package environment

import (
	"bufio"
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/run"
	"strings"
)

func eval(
	path string,
	base []string,
) (map[string]string, error) {
	r := run.New()
	r.Panic = false
	r.SetEnvironment(base)
	output := r.Start(
		"/bin/sh",
		"-c",
		fmt.Sprintf(". %s && env -0", path),
	)

	if r.Error != nil {
		return nil, fmt.Errorf("eval %s: %w", path, r.Error)
	}

	baseSet := make(map[string]string, len(base))

	for _, entry := range base {
		key, value, found := strings.Cut(entry, "=")

		if found {
			baseSet[key] = value
		}
	}

	result := make(map[string]string)
	scanner := bufio.NewScanner(strings.NewReader(output))
	scanner.Split(splitNull)

	for scanner.Scan() {
		line := scanner.Text()
		key, value, found := strings.Cut(line, "=")

		if !found {
			continue
		}

		existing, exists := baseSet[key]

		if !exists || existing != value {
			result[key] = value
		}
	}

	return result, nil
}
