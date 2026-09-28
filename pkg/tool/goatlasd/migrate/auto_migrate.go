package migrate

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/placement"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/model/sighting"
	"gorm.io/gorm"
)

func AutoMigrate(d *gorm.DB) {
	errors.PanicOnError(d.AutoMigrate(placement.Stub()))
	errors.PanicOnError(d.AutoMigrate(sighting.Stub()))
	errors.PanicOnError(d.Exec(constant.PlacementIndex).Error)
	errors.PanicOnError(d.Exec(constant.SightingIndex).Error)
}
