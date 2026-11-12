package generation

import "time"

func Prune(
	root string,
	keep string,
	owns func(string) bool,
) []string {
	return Collect(root, keep, time.Now(), 0, owns)
}
