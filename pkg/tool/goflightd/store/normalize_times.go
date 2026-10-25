package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	relationalConstant "github.com/funtimecoding/soil/pkg/relational/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/event"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/mark"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/snapshot"
	"gorm.io/gorm"
)

func normalizeTimes(d *gorm.DB) {
	if d.Name() != relationalConstant.LiteDialectName {
		return
	}

	for _, table := range []string{
		event.Stub().TableName(),
		snapshot.Stub().TableName(),
		mark.Stub().TableName(),
	} {
		errors.PanicOnError(
			d.Exec(fmt.Sprintf(constant.NormalizeTimeQuery, table)).Error,
		)
	}
}
