package pricing

func Cost(
	model string,
	t *Tokens,
) float64 {
	var inputRate float64
	var outputRate float64

	switch model {
	case "fable", "mythos":
		inputRate = 10.0
		outputRate = 50.0
	case "opus":
		inputRate = 5.0
		outputRate = 25.0
	case "haiku":
		inputRate = 1.0
		outputRate = 5.0
	default:
		inputRate = 3.0
		outputRate = 15.0
	}

	unknown := t.CacheCreation - t.Cache5Minute - t.Cache1Hour

	if unknown < 0 {
		unknown = 0
	}

	fiveMinute := unknown + t.Cache5Minute

	return float64(t.Input)/1_000_000*inputRate +
		float64(t.Output)/1_000_000*outputRate +
		float64(fiveMinute)/1_000_000*inputRate*1.25 +
		float64(t.Cache1Hour)/1_000_000*inputRate*2.0 +
		float64(t.CacheRead)/1_000_000*inputRate*0.1
}
