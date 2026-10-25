package connection

import (
	"database/sql"
	"database/sql/driver"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/relational/constant"
	_ "github.com/glebarez/go-sqlite"
)

func Driver() driver.Driver {
	d, e := sql.Open(constant.LiteDriverName, "")
	errors.PanicOnError(e)
	defer errors.PanicClose(d)

	return d.Driver()
}
