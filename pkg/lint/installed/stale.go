package installed

import "golang.org/x/mod/semver"

func Stale(
	built string,
	latest string,
) bool {
	return semver.Compare(built, latest) < 0
}
