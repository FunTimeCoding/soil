package migrate

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/event"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/mark"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/snapshot"
	"gorm.io/gorm"
)

func AutoMigrate(d *gorm.DB) {
	errors.PanicOnError(
		d.AutoMigrate(event.Stub(), snapshot.Stub(), mark.Stub()),
	)
}
