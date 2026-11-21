package migrate

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/flag"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/installed"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/mirror"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/score"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
	"gorm.io/gorm"
)

func AutoMigrate(d *gorm.DB) {
	migrateFlagSubject(d)
	errors.PanicOnError(
		d.AutoMigrate(
			flag.Stub(),
			system.Stub(),
			installed.Stub(),
			score.Stub(),
			mirror.Stub(),
		),
	)
}
