package migrate

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/model/click"
	"gorm.io/gorm"
)

func AutoMigrate(d *gorm.DB) {
	errors.PanicOnError(d.AutoMigrate(click.New()))
}
