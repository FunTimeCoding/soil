package route_guard

import "regexp"

func inScope(path string) bool {
	return regexp.MustCompile(
		`(^|/)tool/[^/]+/(web|model_context)(/|$)`,
	).MatchString(path)
}
