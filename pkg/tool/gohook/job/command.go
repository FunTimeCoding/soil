package job

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func (j *Job) Command(arguments []string) string {
	result := strings.ReplaceAll(j.Run, "{@}", join.Space(arguments...))

	for i, a := range arguments {
		result = strings.ReplaceAll(result, fmt.Sprintf("{%d}", i+1), a)
	}

	return result
}
