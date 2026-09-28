package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"gorm.io/gorm/clause"
)

func (s *Store) SavePlacement(p *placement.Placement) error {
	return s.mapper.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: constant.SourceColumn},
				{Name: constant.KindColumn},
				{Name: constant.ScopeColumn},
				{Name: constant.NameColumn},
				{Name: constant.PlaceKindColumn},
				{Name: constant.PlaceIdentifierColumn},
			},
			DoUpdates: clause.AssignmentColumns(
				[]string{
					constant.PlaceNameColumn,
					constant.PackageColumn,
					constant.VersionColumn,
					constant.SeenColumn,
				},
			),
		},
	).Create(p).Error
}
