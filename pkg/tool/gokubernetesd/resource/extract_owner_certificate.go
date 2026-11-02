package resource

func ExtractOwnerCertificate(object map[string]any) string {
	metadata, okay := object["metadata"].(map[string]any)

	if !okay {
		return ""
	}

	owners, okay := metadata["ownerReferences"].([]any)

	if !okay {
		return ""
	}

	for _, raw := range owners {
		e, okay := raw.(map[string]any)

		if !okay {
			continue
		}

		if e["kind"] == "Certificate" {
			name, _ := e["name"].(string)

			return name
		}
	}

	return ""
}
