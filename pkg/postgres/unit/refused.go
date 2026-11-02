package unit

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
)

func refused(code string) error {
	return fmt.Errorf(
		"query: %w",
		&pgconn.PgError{Severity: "ERROR", Code: code, Message: "refused"},
	)
}
