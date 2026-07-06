package file_report

import (
	"bufio"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"strings"
)

type Report struct {
	modified        bool
	originalBuilder *strings.Builder
	fixedBuilder    *strings.Builder
	scanner         *bufio.Scanner
	lineNumber      int
	FilePath        string
	Concerns        []*concern.Concern
	Original        string
	Fixed           string
	Fix             func()
}
