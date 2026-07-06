package relational

import "github.com/funtimecoding/soil/pkg/errors"

func (d *Database) Create(a any) {
	errors.PanicOnError(d.mapper.Create(a).Error)
}
