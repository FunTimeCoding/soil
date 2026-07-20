package pricing

func KnownModel(model string) bool {
	switch model {
	case "fable", "mythos", "opus", "sonnet", "haiku":
		return true
	}

	return false
}
