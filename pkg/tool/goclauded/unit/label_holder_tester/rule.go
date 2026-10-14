package label_holder_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/label_holder"
	"time"
)

func Rule() *label_holder.Rule {
	return label_holder.NewRule("role", "watch", 30*time.Minute, 5)
}
