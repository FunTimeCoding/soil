package decision_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
)

func Mark(identifier uint) string {
	return fmt.Sprintf("%s%d", constant.DecisionMarkPrefix, identifier)
}
