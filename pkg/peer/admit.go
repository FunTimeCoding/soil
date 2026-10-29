package peer

func (g *Guard) admit() (bool, error) {
	e := g.Unavailable()

	if e == nil {
		return false, nil
	}

	g.mutex.Lock()
	defer g.mutex.Unlock()

	if !g.served || g.trying || g.clock().Before(g.retry) {
		return false, e
	}

	g.trying = true

	return true, nil
}
