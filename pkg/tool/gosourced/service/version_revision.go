package service

import "golang.org/x/mod/module"

func versionRevision(version string) string {
	if !module.IsPseudoVersion(version) {
		return version
	}

	result, e := module.PseudoVersionRev(version)

	if e != nil {
		return version
	}

	return result
}
