package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func legacyMetadata(
	kind string,
	scope string,
	body string,
) map[string]string {
	switch kind {
	case "announce":
		if body != "" {
			return map[string]string{constant.Topic: body}
		}
	case "complete", "update":
		result := map[string]string{}

		if scope != "" {
			result[constant.Topic] = scope
		}

		if body != "" {
			result[constant.Message] = body
		}

		if len(result) > 0 {
			return result
		}
	case "moment":
		if body != "" {
			return map[string]string{constant.Line: body}
		}
	case "summarize":
		if body != "" {
			return map[string]string{constant.Body: body}
		}
	case "label":
		result := map[string]string{}

		if scope != "" {
			result[constant.Target] = scope
		}

		if body != "" {
			result["change"] = body
		}

		if len(result) > 0 {
			return result
		}
	case "inactivity_timeout":
		if scope != "" {
			return map[string]string{constant.Topic: scope}
		}
	case "send":
		result := map[string]string{}

		if scope != "" {
			result["recipient"] = scope
		}

		if body != "" {
			result[constant.Message] = body
		}

		if len(result) > 0 {
			return result
		}
	}

	return nil
}
