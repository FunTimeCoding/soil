package brave

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/brave/constant"
	"github.com/funtimecoding/soil/pkg/brave/helper"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func OpenProfile(profile string) {
	run.New().Launch(
		constant.BravePath,
		fmt.Sprintf("--user-data-dir=%s", helper.SettingsPath()),
		fmt.Sprintf("--profile-directory=%s", profile),
	)
}
