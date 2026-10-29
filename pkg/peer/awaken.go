package peer

func (g *Guard) awaken() {
	g.mutex.Lock()
	served := g.served
	g.mutex.Unlock()

	if served {
		return
	}

	select {
	case g.wake <- struct{}{}:
	default:
	}
}
