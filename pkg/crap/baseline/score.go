package baseline

func (b *Baseline) Score(key string) (float64, bool) {
	v, okay := b.scores[key]

	return v, okay
}
