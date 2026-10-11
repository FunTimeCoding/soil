package release

import "github.com/coreos/go-semver/semver"

type Release struct {
	Codename string
	Major    int64
	Minor    int64
	version  *semver.Version
}
