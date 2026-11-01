package component

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/constant"
)

func PlayerPath(account string) string {
	return fmt.Sprintf("%s/%s", constant.PlayersPath, account)
}
