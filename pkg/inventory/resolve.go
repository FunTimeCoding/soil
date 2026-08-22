package inventory

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"slices"
)

func Resolve(
	explicit string,
	names []string,
) (string, error) {
	if explicit != "" {
		if !slices.Contains(names, explicit) {
			return "", not_found.New("instance", explicit)
		}

		return explicit, nil
	}

	if len(names) == 1 {
		return names[0], nil
	}

	return "", not_selected.Format(
		"no instance selected - %d instances configured, selection required",
		len(names),
	)
}
