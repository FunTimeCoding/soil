package label_holder_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/label_holder"
	"time"
)

func Rule(compatible ...string) *label_holder.Rule {
	return label_holder.NewRule(
		constant.FixtureLabelKey,
		"watch",
		compatible,
		30*time.Minute,
		5,
	)
}
