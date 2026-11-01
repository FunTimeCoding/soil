package hold

import "time"

func (h *Hold) loop() {
	delay := h.minimum

	for h.context.Err() == nil {
		served := false
		h.recovery.Run(func() { served = h.session() })

		if served {
			delay = h.minimum
		}

		select {
		case <-h.context.Done():
			return
		case <-time.After(delay):
		}

		if !served {
			delay = min(delay*2, h.maximum)
		}
	}
}
