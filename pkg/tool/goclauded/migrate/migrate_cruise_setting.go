package migrate

import (
	"github.com/funtimecoding/soil/pkg/relational/schema"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"gorm.io/gorm"
)

func migrateCruiseSetting(d *gorm.DB) {
	schema.DropColumns(
		d,
		constant.CruiseSettingTable,
		constant.CruiseModeColumn,
		"pace",
		"promoted_at",
	)
}
