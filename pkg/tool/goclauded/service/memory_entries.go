package service

import "encoding/json"

func memoryEntries(body string) []memoryEntry {
	var many []memoryEntry

	if json.Unmarshal([]byte(body), &many) == nil {
		return many
	}

	var one memoryEntry

	if json.Unmarshal([]byte(body), &one) == nil && one.Identifier > 0 {
		return []memoryEntry{one}
	}

	return nil
}
