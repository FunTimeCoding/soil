package migrate

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/provision/model/run"
	"gorm.io/gorm"
)

func AutoMigrate(
	d *gorm.DB,
	tableName string,
) {
	errors.PanicOnError(d.Table(tableName).AutoMigrate(run.New()))
}
