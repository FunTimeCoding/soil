package peer

func (g *Guard) Close() {
	g.cancel()
}
