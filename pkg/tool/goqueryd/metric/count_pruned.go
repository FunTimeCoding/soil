package metric

func (m *Metric) CountPruned(contents int64) {
	m.pruned.Add(float64(contents))
}
