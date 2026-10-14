package label_holder_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/target"
	"time"
)

func Idle(
	identifier string,
	role string,
) *target.Target {
	return Seen(identifier, role, 2*time.Hour, 10)
}
