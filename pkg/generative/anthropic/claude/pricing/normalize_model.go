package pricing

import "strings"

func NormalizeModel(model string) string {
	lower := strings.ToLower(model)

	if strings.Contains(lower, "fable") {
		return "fable"
	}

	if strings.Contains(lower, "mythos") {
		return "mythos"
	}

	if strings.Contains(lower, "opus") {
		return "opus"
	}

	if strings.Contains(lower, "sonnet") {
		return "sonnet"
	}

	if strings.Contains(lower, "haiku") {
		return "haiku"
	}

	return model
}
