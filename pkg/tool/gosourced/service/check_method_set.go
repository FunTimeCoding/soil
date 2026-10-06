package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/relocation"
	"go/types"
)

func checkMethodSet(entries []*relocation.Entry) string {
	for _, entry := range entries {
		if _, okay := entry.Object.(*types.TypeName); !okay {
			continue
		}

		named, okay := entry.Object.Type().(*types.Named)

		if !okay || named.NumMethods() == 0 {
			continue
		}

		return fmt.Sprintf(
			"%s has methods - use %s to move a type with its method set",
			entry.Symbol,
			constant.ExtractType,
		)
	}

	return ""
}
