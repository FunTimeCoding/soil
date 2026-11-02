package runtime

import (
	"github.com/coreos/go-semver/semver"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/expression"
	"github.com/funtimecoding/soil/pkg/system"
)

func ExecutableVersion() *semver.Version {
	if m := expression.SubMatch(
		`go(\d+\.\d+\.\d+)`,
		system.Run(soil.Go, soil.Version),
	); m != "" {
		return semver.New(m)
	}

	return nil
}
