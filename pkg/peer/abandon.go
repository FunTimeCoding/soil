package peer

func (g *Guard) abandon(trial bool) {
	if !trial {
		return
	}

	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.trying = false
}
