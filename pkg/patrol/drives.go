package patrol

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/smartctl"
	smartctlConstant "github.com/funtimecoding/soil/pkg/smartctl/constant"
	"github.com/funtimecoding/soil/pkg/smartctl/response"
)

func (c *Client) Drives(_ context.Context) ([]*response.Report, error) {
	s, e := c.dial()

	if e != nil {
		return nil, e
	}

	defer s.Close()
	r := s.Run(smartctlConstant.ScanCommand)

	if r.Error != nil {
		return nil, r.Error
	}

	var scan response.Scan

	if f := notation.Decode(r.OutputString, &scan); f != nil {
		return nil, fmt.Errorf("decode scan: %w", f)
	}

	var result []*response.Report

	for _, d := range scan.Devices {
		if !smartctlConstant.DevicePath.MatchString(d.Name) {
			return nil, unexpected.Format("unexpected device: %s", d.Name)
		}

		v := s.Run(fmt.Sprintf(smartctlConstant.ReportCommand, d.Name))
		report, f := smartctl.Report(v.OutputString)

		if f != nil {
			if v.Error != nil {
				return nil, v.Error
			}

			return nil, f
		}

		result = append(result, report)
	}

	return result, nil
}
