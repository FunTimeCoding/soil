package error_answer

import "regexp"

func inScope(path string) bool {
	return regexp.MustCompile(
		`(^|/)tool/[^/]+/(server|model_context)(/|$)`,
	).MatchString(path)
}
