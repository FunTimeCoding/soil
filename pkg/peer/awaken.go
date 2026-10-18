package peer

func (g *Guard) awaken() {
	select {
	case g.wake <- struct{}{}:
	default:
	}
}
