package service

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func profileEntries(body string) map[string][]memoryEntry {
	var payload profilePayload

	if json.Unmarshal([]byte(body), &payload) != nil {
		return nil
	}

	return map[string][]memoryEntry{
		constant.TierAlways:   payload.Always,
		constant.TierRelevant: payload.Relevant,
	}
}
