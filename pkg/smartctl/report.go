package smartctl

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/smartctl/constant"
	"github.com/funtimecoding/soil/pkg/smartctl/response"
)

func Report(output string) (*response.Report, error) {
	var result response.Report

	if e := notation.Decode(output, &result); e != nil {
		return nil, fmt.Errorf("decode report: %w", e)
	}

	if result.Run == nil || result.Device == nil {
		return nil, unexpected.Format("unexpected report: %s", output)
	}

	if result.Run.ExitStatus&constant.UnreadMask != 0 {
		return nil, unexpected.Format(
			"unread device %s: exit status %d",
			result.Device.Name,
			result.Run.ExitStatus,
		)
	}

	return &result, nil
}
