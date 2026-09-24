package entry

func (e *Entry) KillRatio() float64 {
	total := e.Killed + len(e.Lived)

	if total == 0 {
		return 1
	}

	return float64(e.Killed) / float64(total)
}
