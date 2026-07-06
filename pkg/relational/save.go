package relational

import "github.com/funtimecoding/soil/pkg/errors"

func (d *Database) Save(a any) {
	errors.PanicOnError(d.mapper.Save(a).Error)
}
