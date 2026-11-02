package system

import (
	"fmt"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"os"
	"strings"
)

func PrintEnvironment() {
	for _, e := range os.Environ() {
		p := strings.SplitN(e, stringsConstant.Equals, 2)
		fmt.Printf("ENV: %s\n", key_value.Equals(p[0], p[1]))
	}
}
