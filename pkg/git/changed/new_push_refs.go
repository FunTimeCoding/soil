package changed

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/strings/slice"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"strings"
)

func NewPushRefs(
	directory string,
	input string,
) *Range {
	lines := slice.StripEmpty(split.NewLine(strings.TrimSpace(input)))

	if len(lines) == 0 {
		return NewPush(directory)
	}

	for _, line := range lines {
		fields := strings.Fields(line)

		if len(fields) < 4 {
			continue
		}

		local, remote := fields[1], fields[3]

		if local == constant.ZeroHash {
			continue
		}

		if remote == constant.ZeroHash {
			return NewPush(directory)
		}

		return New(remote, local)
	}

	return NewNone()
}
