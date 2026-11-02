package resource

import "fmt"

func StripArgocdNoise(
	object map[string]any,
	filtered *[]string,
) {
	status, okay := object["status"].(map[string]any)

	if !okay {
		return
	}

	if resources, okay := status["resources"].([]any); okay {
		*filtered = append(
			*filtered,
			fmt.Sprintf(
				"status.resources (%d entries, all synced)",
				len(resources),
			),
		)
		delete(status, "resources")
	}

	if history, okay := status["history"].([]any); okay {
		*filtered = append(
			*filtered,
			fmt.Sprintf("status.history (%d entries)", len(history)),
		)
		delete(status, "history")
	}

	if result, okay := status["operationState"].(map[string]any); okay {
		if syncResult, okay := result["syncResult"].(map[string]any); okay {
			if resources, okay := syncResult["resources"].([]any); okay {
				*filtered = append(
					*filtered,
					fmt.Sprintf(
						"status.operationState.syncResult.resources (%d entries)",
						len(resources),
					),
				)
				delete(syncResult, "resources")
			}
		}
	}
}
