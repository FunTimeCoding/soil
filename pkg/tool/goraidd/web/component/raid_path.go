package component

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
)

func RaidPath(identifier uint) string {
	return fmt.Sprintf("%s/%d", constant.RaidsPath, identifier)
}
