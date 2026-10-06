package service

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/memory_payload"
)

func profileEntries(body string) map[string][]memory_payload.Entry {
	var payload memory_payload.Profile

	if json.Unmarshal([]byte(body), &payload) != nil {
		return nil
	}

	return map[string][]memory_payload.Entry{
		constant.TierAlways:   payload.Always,
		constant.TierRelevant: payload.Relevant,
	}
}
