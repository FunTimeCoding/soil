package service

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/memory_payload"
)

func groupEntries(body string) []memory_payload.Entry {
	var payload memory_payload.Group

	if json.Unmarshal([]byte(body), &payload) != nil {
		return nil
	}

	if payload.Parent == nil || payload.Parent.Identifier == 0 {
		return nil
	}

	return append([]memory_payload.Entry{*payload.Parent}, payload.Children...)
}
