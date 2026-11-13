package relational

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"gorm.io/gorm"
)

func FetchOne(
	q *gorm.DB,
	out any,
	kind string,
	identifier any,
) error {
	found, e := FindOne(q, out)

	if e != nil {
		return e
	}

	if !found {
		return not_found.New(kind, identifier)
	}

	return nil
}
