package peer

func (g *Guard) loop() {
	for {
		select {
		case <-g.context.Done():
			return
		case <-g.wake:
		}

		g.restore()
	}
}
