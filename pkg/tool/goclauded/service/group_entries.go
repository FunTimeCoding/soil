package service

import "encoding/json"

func groupEntries(body string) []memoryEntry {
	var payload groupPayload

	if json.Unmarshal([]byte(body), &payload) != nil {
		return nil
	}

	if payload.Parent == nil || payload.Parent.Identifier == 0 {
		return nil
	}

	return append([]memoryEntry{*payload.Parent}, payload.Children...)
}
