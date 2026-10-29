package event

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/event"
	"time"
)

func deleteOld(
	c *client.Client,
	v *event.Event,
	age time.Duration,
) bool {
	if v.Age() > age {
		c.MustDeleteEventSimple(v)

		return true
	}

	return false
}
