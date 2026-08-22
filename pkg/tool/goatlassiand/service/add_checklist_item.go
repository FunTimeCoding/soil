package service

import "github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/checklist_item"

func (s *Service) AddChecklistItem(
	key string,
	text string,
) ([]*checklist_item.Item, error) {
	items, e := s.ReadChecklist(key)

	if e != nil {
		return nil, e
	}

	items = append(items, checklist_item.New(len(items), text, false))

	if f := s.WriteChecklist(key, items); f != nil {
		return nil, f
	}

	return items, nil
}
