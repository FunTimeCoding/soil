package lite

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/relational/constant"
	"github.com/glebarez/go-sqlite"
	"modernc.org/sqlite/lib"
)

func IsBusy(e error) bool {
	var f *sqlite.Error

	if errors.As(e, &f) {
		return f.Code()&constant.LitePrimaryCodeMask == sqlite3.SQLITE_BUSY
	}

	return false
}
