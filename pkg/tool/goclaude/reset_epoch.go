package goclaude

func resetEpoch(v int64) *int64 {
	if v == 0 {
		return nil
	}

	return new(v)
}
