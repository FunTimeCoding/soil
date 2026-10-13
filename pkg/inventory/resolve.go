package inventory

import (
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/inventory/constant"
	"slices"
)

func Resolve(
	explicit string,
	names []string,
) (string, error) {
	if explicit != "" {
		if !slices.Contains(names, explicit) {
			return "", not_found.New(constant.InstanceKind, explicit)
		}

		return explicit, nil
	}

	if len(names) == 0 {
		return "", not_configured.Format("no instance configured")
	}

	if len(names) == 1 {
		return names[0], nil
	}

	return "", not_selected.Format(
		"no instance selected - %d instances configured, selection required",
		len(names),
	)
}
