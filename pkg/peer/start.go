package peer

func (g *Guard) Start() {
	go g.loop()
}
