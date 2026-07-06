package event

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/event"
	"slices"
)

func deleteIrrelevant(
	c *client.Client,
	e *event.Event,
) bool {
	if slices.Contains(constant.IrrelevantEventReason, e.Reason) {
		c.DeleteEventSimple(e)

		return true
	}

	return false
}
