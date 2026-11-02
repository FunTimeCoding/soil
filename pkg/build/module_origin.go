package build

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/build/origin"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func ModuleOrigin(
	module string,
	version string,
) (*origin.Origin, error) {
	if version == "" {
		version = soil.LatestVersion
	}

	r := run.New()
	r.Panic = false
	output := r.Start(
		soil.Go,
		soil.List,
		soil.ModuleArgument,
		soil.NotationArgument,
		fmt.Sprintf("%s@%s", module, version),
	)

	if r.Error != nil {
		return nil, r.Error
	}

	if v := ParseOrigin(output); v != nil && v.Version != "" {
		return v, nil
	}

	return nil, fmt.Errorf("no version for %s@%s", module, version)
}
