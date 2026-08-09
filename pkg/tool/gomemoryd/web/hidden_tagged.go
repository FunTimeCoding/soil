package web

import "slices"

func hiddenTagged(
	tags []string,
	tag string,
) bool {
	return tag != "" && slices.Contains(tags, tag)
}
