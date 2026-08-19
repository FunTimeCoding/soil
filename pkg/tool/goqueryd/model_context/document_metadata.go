package model_context

func documentMetadata(scalar map[string]string) map[string][]string {
	if len(scalar) == 0 {
		return nil
	}

	result := map[string][]string{}

	for key, value := range scalar {
		result[key] = []string{value}
	}

	return result
}
