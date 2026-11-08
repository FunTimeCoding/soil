package label_holder_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/target"
	"time"
)

func Seen(
	identifier string,
	role string,
	ago time.Duration,
	turns int,
) *target.Target {
	result := target.New(
		identifier,
		identifier,
		Now().Add(-ago),
		map[string]string{constant.FixtureLabelKey: role},
	)
	result.TurnCount = turns

	return result
}
