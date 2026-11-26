package service

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/types/node"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/types/view"
)

func (s *Service) observe(tab string) (*view.View, error) {
	settled := s.navigation.Settle(tab)
	raw, e := guarded(
		s,
		func() (json.RawMessage, error) {
			return s.client.CallFunction(tab, constant.SnapshotScript)
		},
	)

	if e != nil {
		return nil, e
	}

	tree, shared, f := decodeSnapshot(raw)

	if f != nil {
		return nil, f
	}

	references := make(map[string]string)
	root := node.New("", constant.DocumentRole, tree.Title, "")
	root.Children = convertItems(tree.Items, shared, references)
	s.mutex.Lock()
	s.references[tab] = references
	s.mutex.Unlock()
	result, g := s.present(tab, node.Format([]*node.Node{root}, 0))

	if g != nil {
		return nil, g
	}

	result.Loading = !settled

	return result, nil
}
