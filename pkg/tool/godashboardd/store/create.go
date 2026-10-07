package store

import "github.com/funtimecoding/soil/pkg/tool/godashboardd/model/click"

func (s *Store) Create(label string) error {
	c := click.New()
	c.Label = label

	return s.mapper.Create(c).Error
}
