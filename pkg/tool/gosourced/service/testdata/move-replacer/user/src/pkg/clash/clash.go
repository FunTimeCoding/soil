package clash

import (
	"example/pkg/target"
	"other.test/lib/source"
)

func Use() string {
	return target.Local() + source.Helper()
}
