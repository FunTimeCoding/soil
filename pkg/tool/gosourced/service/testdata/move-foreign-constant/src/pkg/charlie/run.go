package charlie

import (
	"example/pkg/target"
	"fmt"
)

func Run() string {
	return fmt.Sprint(target.Mode)
}
