package netbox

func localContext(payload map[string]any) any {
	if len(payload) == 0 {
		return map[string]any(nil)
	}

	return payload
}
