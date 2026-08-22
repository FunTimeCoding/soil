package service

import "github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/checklist_item"

func (s *Service) DeleteChecklistItem(
	key string,
	index int,
) ([]*checklist_item.Item, error) {
	items, e := s.ReadChecklist(key)

	if e != nil {
		return nil, e
	}

	if f := checklistAt(items, index); f != nil {
		return nil, f
	}

	items = append(items[:index-1], items[index:]...)

	for j := range items {
		items[j].Index = j + 1
	}

	if f := s.WriteChecklist(key, items); f != nil {
		return nil, f
	}

	return items, nil
}
