package coverage

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/system/run"
	"path/filepath"
)

func RunTestAlone(
	root string,
	binary string,
	test string,
	directory string,
) string {
	profile := filepath.Join(
		directory,
		join.Empty(test, constant.ProfileSuffix),
	)
	r := run.New()
	r.Panic = false
	r.Directory = root
	r.Start(
		binary,
		key_value.Equals(
			constant.TestRun,
			fmt.Sprintf(constant.ExactPattern, test),
		),
		key_value.Equals(constant.TestCoverProfile, profile),
	)

	return profile
}
