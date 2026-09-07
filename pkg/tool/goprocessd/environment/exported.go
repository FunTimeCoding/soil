package environment

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func exported(
	path string,
	base []string,
) (map[string]struct{}, error) {
	minimal := harness(base)
	baselineRun := run.New()
	baselineRun.Panic = false
	baselineRun.SetEnvironment(minimal)
	baselineOutput := baselineRun.Start("/bin/sh", "-c", "env -0")

	if baselineRun.Error != nil {
		return nil, fmt.Errorf("baseline: %w", baselineRun.Error)
	}

	sourcedRun := run.New()
	sourcedRun.Panic = false
	sourcedRun.SetEnvironment(minimal)
	sourcedOutput := sourcedRun.Start(
		"/bin/sh",
		"-c",
		fmt.Sprintf(". %s && env -0", path),
	)

	if sourcedRun.Error != nil {
		return nil, fmt.Errorf("exported %s: %w", path, sourcedRun.Error)
	}

	baseline := parseNull(baselineOutput)
	result := make(map[string]struct{})

	for key, value := range parseNull(sourcedOutput) {
		previous, present := baseline[key]

		if !present || previous != value {
			result[key] = struct{}{}
		}
	}

	return result, nil
}
