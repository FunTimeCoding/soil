package resource

func ExtractDnsNames(object map[string]any) []string {
	raw, okay := object["spec"].(map[string]any)["dnsNames"]

	if !okay {
		return nil
	}

	list, okay := raw.([]any)

	if !okay {
		return nil
	}

	var result []string

	for _, item := range list {
		if s, okay := item.(string); okay {
			result = append(result, s)
		}
	}

	return result
}
