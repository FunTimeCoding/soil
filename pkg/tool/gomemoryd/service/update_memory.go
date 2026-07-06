package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
)

func (s *Service) UpdateMemory(
	identifier int64,
	name string,
	content string,
	description string,
	source string,
) (*store.Memory, error) {
	existing, e := s.store.GetMemory(identifier)

	if e != nil {
		return nil, e
	}

	o := save_option.New()
	o.Name = name
	o.Content = content
	o.Description = description
	o.Tags = existing.Tags
	o.Source = source
	e = s.store.UpdateMemory(identifier, o)

	if e != nil {
		return nil, e
	}

	m, e := s.store.GetMemory(identifier)

	if e != nil {
		return nil, e
	}

	if e = s.indexer.Push(
		fmt.Sprintf("memory/%d", m.Identifier),
		m.Content,
		map[string]string{
			"memory_id":   fmt.Sprintf("%d", m.Identifier),
			"type":        m.Type,
			"name":        m.Name,
			"description": m.Description,
		},
	); e != nil {
		return nil, e
	}

	return m, nil
}
