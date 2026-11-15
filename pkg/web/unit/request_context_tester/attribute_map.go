package request_context_tester

import "log/slog"

func AttributeMap(attributes []slog.Attr) map[string]any {
	result := map[string]any{}

	for _, a := range attributes {
		result[a.Key] = a.Value.Any()
	}

	return result
}
