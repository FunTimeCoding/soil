package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/sighting"
	"gorm.io/gorm/clause"
)

func (s *Store) SaveSighting(v *sighting.Sighting) error {
	return s.mapper.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: constant.SourceColumn},
				{Name: constant.HardwareAddressColumn},
			},
			DoUpdates: clause.AssignmentColumns(
				[]string{
					constant.AddressColumn,
					constant.HostnameColumn,
					constant.ReservedColumn,
					constant.PlaceKindColumn,
					constant.PlaceIdentifierColumn,
					constant.PlaceNameColumn,
					constant.SeenColumn,
				},
			),
		},
	).Create(v).Error
}
