package label_holder_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/target"
	"time"
)

func Seen(
	identifier string,
	role string,
	ago time.Duration,
	turns int,
) *target.Target {
	labels := map[string]string{}

	if role != "" {
		labels["role"] = role
	}

	result := target.New(identifier, identifier, Now().Add(-ago), labels)
	result.TurnCount = turns

	return result
}
