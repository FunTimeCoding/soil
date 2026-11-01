package hold

import (
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/peer/constant"
	"time"
)

func (h *Hold) Unavailable() error {
	h.mutex.Lock()
	serving := h.serving
	h.mutex.Unlock()

	if serving {
		return nil
	}

	if down, x := h.reacher.Down(h.name); down {
		return unavailable.New(h.name, x.Since, x.Reason)
	}

	return unavailable.New(h.name, time.Now(), constant.NotConnectedReason)
}
