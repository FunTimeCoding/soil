package system

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"os"
	"strings"
)

func PrintEnvironment() {
	for _, e := range os.Environ() {
		p := strings.SplitN(e, separator.Equals, 2)
		fmt.Printf("ENV: %s\n", key_value.Equals(p[0], p[1]))
	}
}
