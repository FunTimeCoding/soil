package outage_tester

func placeholder(p property) any {
	if len(p.Enum) > 0 {
		return p.Enum[0]
	}

	switch p.Type {
	case "integer", "number":
		return 1
	case "boolean":
		return false
	case "array":
		return []any{"1"}
	case "object":
		return map[string]any{}
	default:
		return "1"
	}
}
