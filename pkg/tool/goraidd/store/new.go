package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/raid/model/fight"
	"github.com/funtimecoding/soil/pkg/raid/model/player_fight_statistic"
	"github.com/funtimecoding/soil/pkg/raid/model/raid"
	"gorm.io/gorm"
)

func New(m *gorm.DB) *Store {
	errors.PanicOnError(m.AutoMigrate(raid.New()))
	errors.PanicOnError(m.AutoMigrate(fight.New()))
	errors.PanicOnError(m.AutoMigrate(player_fight_statistic.New()))

	return &Store{mapper: m}
}
