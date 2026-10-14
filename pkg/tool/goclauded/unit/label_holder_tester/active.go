package label_holder_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/target"
	"time"
)

func Active(
	identifier string,
	role string,
) *target.Target {
	return Seen(identifier, role, time.Minute, 10)
}
