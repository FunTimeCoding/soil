package user

import (
	"context"
	"other.test/lib"
)

func Use(x context.Context) string {
	return lib.Log("a")
}
