package pod

func (p *Pod) WithRestarts(restarts int64) *Pod {
	p.Restarts = restarts

	return p
}
