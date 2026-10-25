package nil_dependency

import "strings"

func isDependencyPath(path string) bool {
	if strings.HasSuffix(path, "/pkg/face") {
		return true
	}

	_, tool, found := strings.Cut(path, "/pkg/tool/")

	if !found {
		return false
	}

	segments := strings.Split(tool, "/")

	return len(segments) == 2 && segments[1] == "face"
}
