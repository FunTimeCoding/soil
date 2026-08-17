package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/types"
)

func checkMethodSet(entries []*moveEntry) string {
	for _, entry := range entries {
		if _, okay := entry.object.(*types.TypeName); !okay {
			continue
		}

		named, okay := entry.object.Type().(*types.Named)

		if !okay || named.NumMethods() == 0 {
			continue
		}

		return fmt.Sprintf(
			"%s has methods - use %s to move a type with its method set",
			entry.symbol,
			constant.ExtractType,
		)
	}

	return ""
}
