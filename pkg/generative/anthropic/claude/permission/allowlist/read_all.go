package allowlist

import (
	"errors"
	"io/fs"
)

func ReadAll(paths ...string) ([]*Rule, error) {
	var result []*Rule

	for _, path := range paths {
		rules, e := Read(path)

		if errors.Is(e, fs.ErrNotExist) {
			continue
		}

		if e != nil {
			return nil, e
		}

		result = append(result, rules...)
	}

	return result, nil
}
