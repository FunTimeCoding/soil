package service

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/browser/types/response"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/types/page_tree"
)

func decodeSnapshot(raw json.RawMessage) (*page_tree.Tree, []string, error) {
	outer := response.NewRemoteValue()

	if e := notation.DecodeBytes(raw, outer); e != nil {
		return nil, nil, fmt.Errorf("decode snapshot: %w", e)
	}

	var items []*response.RemoteValue

	if e := notation.DecodeBytes(outer.Value, &items); e != nil {
		return nil, nil, fmt.Errorf("decode snapshot items: %w", e)
	}

	if len(items) == 0 {
		return nil, nil, unexpected.Format("snapshot script answered nothing")
	}

	var payload string

	if e := notation.DecodeBytes(items[0].Value, &payload); e != nil {
		return nil, nil, fmt.Errorf("decode snapshot payload: %w", e)
	}

	result := page_tree.NewTree()

	if e := notation.DecodeBytes([]byte(payload), result); e != nil {
		return nil, nil, fmt.Errorf("decode snapshot tree: %w", e)
	}

	var shared []string

	for _, v := range items[1:] {
		shared = append(shared, v.SharedIdentifier)
	}

	return result, shared, nil
}
