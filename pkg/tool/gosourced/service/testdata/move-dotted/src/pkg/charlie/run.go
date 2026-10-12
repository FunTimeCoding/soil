package charlie

import (
	"example.org/m/pkg/target"
	"fmt"
)

func Run() string {
	return fmt.Sprint(target.Mode)
}
