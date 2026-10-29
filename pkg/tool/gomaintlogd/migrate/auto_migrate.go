package migrate

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"gorm.io/gorm"
)

func AutoMigrate(d *gorm.DB) {
	errors.PanicOnError(d.AutoMigrate(entry.New()))
}
