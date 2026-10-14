package label_holder_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/label_holder"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/target"
)

func Choose(targets ...*target.Target) *label_holder.Plan {
	return label_holder.Choose(targets, Rule(), Now())
}
