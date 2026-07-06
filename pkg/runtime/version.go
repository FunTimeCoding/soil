package runtime

import (
	"github.com/coreos/go-semver/semver"
	"github.com/funtimecoding/soil/pkg/constant"
	"runtime"
	"strings"
)

func Version() *semver.Version {
	return semver.New(strings.TrimPrefix(runtime.Version(), constant.Go))
}
