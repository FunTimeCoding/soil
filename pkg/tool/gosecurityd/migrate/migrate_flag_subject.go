package migrate

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/schema"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/flag"
	"gorm.io/gorm"
)

func migrateFlagSubject(d *gorm.DB) {
	schema.RenameColumns(
		d,
		flag.Stub(),
		map[string]string{
			constant.FlagEarlierColumn: constant.FlagSubjectColumn,
		},
	)

	if m := d.Migrator(); m.HasIndex(flag.Stub(), constant.FlagEarlierIndex) {
		errors.PanicOnError(m.DropIndex(flag.Stub(), constant.FlagEarlierIndex))
	}
}
