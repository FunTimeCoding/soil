package git

import (
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"strconv"
	"strings"
	"time"
)

func CommitTimes(directory string) (map[string]time.Time, error) {
	r := run.New().NoPanic()
	r.Directory = directory
	r.Start(
		constant.Command,
		constant.Log,
		constant.CommitTimeFormat,
		constant.NameOnly,
		constant.Relative,
		constant.Pathspec,
		library.CurrentDirectory,
	)

	if r.Error != nil {
		return nil, r.Error
	}

	result := make(map[string]time.Time)
	var current time.Time

	for _, line := range strings.Split(r.OutputString, "\n") {
		if line == "" {
			continue
		}

		if n, e := strconv.ParseInt(line, 10, 64); e == nil {
			current = time.Unix(n, 0)

			continue
		}

		if _, okay := result[line]; !okay {
			result[line] = current
		}
	}

	return result, nil
}
