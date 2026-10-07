package model_context

import "github.com/funtimecoding/soil/pkg/tool/goclauded/types/caller"

func displayName(
	c *caller.Caller,
	target string,
) string {
	if target != "" {
		return target
	}

	return c.Callsign
}
