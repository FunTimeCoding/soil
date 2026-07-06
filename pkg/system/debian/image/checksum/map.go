package checksum

import (
	"github.com/funtimecoding/soil/pkg/system"
	"strings"
)

func Map(workDirectory string) map[string]string {
	return Parse(strings.TrimSpace(system.ReadFile(workDirectory, File)))
}
