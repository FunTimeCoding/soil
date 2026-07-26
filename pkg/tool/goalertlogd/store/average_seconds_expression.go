package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/relational/constant"
	"gorm.io/gorm"
)

func averageSecondsExpression(d *gorm.DB) string {
	switch d.Name() {
	case constant.LiteDialectName:
		return "(julianday(ended_at) - julianday(started_at)) * 86400.0"
	case constant.PostgresDialectName:
		return "EXTRACT(EPOCH FROM (ended_at - started_at))"
	default:
		panic(fmt.Sprintf("unsupported dialect: %s", d.Name()))
	}
}
