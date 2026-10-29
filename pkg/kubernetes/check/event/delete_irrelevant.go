package event

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/event"
	"slices"
)

func deleteIrrelevant(
	c *client.Client,
	v *event.Event,
) bool {
	if slices.Contains(constant.IrrelevantEventReason, v.Reason) {
		c.MustDeleteEventSimple(v)

		return true
	}

	return false
}
