package statement

import "database/sql/driver"

type Inner interface {
	driver.Stmt
	driver.StmtExecContext
	driver.StmtQueryContext
}
