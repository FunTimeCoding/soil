package watcher

func (w *Watcher) listen() {
	s := w.client.WebSocket()
	s.Listen()

	for {
		select {
		case <-w.done:
			s.Close()

			return
		case v, okay := <-s.EventChannel:
			if !okay {
				if !w.reconnect("event channel closed") {
					return
				}

				s = w.client.WebSocket()

				continue
			}

			w.Dispatch(v)
		case <-s.PingTimeoutChannel:
			if !w.reconnect("ping timeout") {
				return
			}

			s = w.client.WebSocket()
		}
	}
}
