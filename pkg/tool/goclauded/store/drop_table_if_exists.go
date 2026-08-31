package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"gorm.io/gorm"
)

func dropTableIfExists(
	d *gorm.DB,
	table string,
) {
	errors.PanicOnError(
		d.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table)).Error,
	)
}
