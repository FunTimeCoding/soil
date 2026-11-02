package resource

func ExtractConditionStatus(object map[string]any) string {
	status, okay := object["status"].(map[string]any)

	if !okay {
		return "Unknown"
	}

	conditions, okay := status["conditions"].([]any)

	if !okay {
		return "Unknown"
	}

	for _, raw := range conditions {
		condition, okay := raw.(map[string]any)

		if !okay {
			continue
		}

		if condition["type"] == "Ready" {
			if condition["status"] == "True" {
				return "Ready"
			}

			reason, _ := condition["reason"].(string)

			if reason != "" {
				return reason
			}

			return "NotReady"
		}
	}

	return "Unknown"
}
