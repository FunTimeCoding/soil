package pod

func (p *Pod) WithLabels(labels map[string]string) *Pod {
	p.Labels = labels

	return p
}
