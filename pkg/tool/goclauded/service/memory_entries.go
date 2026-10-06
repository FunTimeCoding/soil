package service

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/memory_payload"
)

func memoryEntries(body string) []memory_payload.Entry {
	var many []memory_payload.Entry

	if json.Unmarshal([]byte(body), &many) == nil && len(many) > 0 {
		return many
	}

	if group := groupEntries(body); len(group) > 0 {
		return group
	}

	var one memory_payload.Entry

	if json.Unmarshal([]byte(body), &one) == nil && one.Identifier > 0 {
		return []memory_payload.Entry{one}
	}

	return nil
}
