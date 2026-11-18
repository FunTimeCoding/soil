package nut

import (
	"bufio"
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/nut/constant"
	"github.com/funtimecoding/soil/pkg/nut/detached"
	"io"
	"net"
	"strings"
	"time"
)

func (c *Client) exchange(
	x context.Context,
	command string,
	subject string,
) ([]string, error) {
	d := net.Dialer{Timeout: c.timeout}
	n, e := d.DialContext(x, constant.Network, c.address)

	if e != nil {
		return nil, e
	}

	defer errors.LogClose(n)

	if e = n.SetDeadline(time.Now().Add(c.timeout)); e != nil {
		return nil, e
	}

	if _, e = fmt.Fprintf(
		n,
		constant.ExchangeFormat,
		command,
		constant.Logout,
	); e != nil {
		return nil, e
	}

	s := bufio.NewScanner(n)
	var result []string

	for s.Scan() {
		line := s.Text()

		if code, found := strings.CutPrefix(line, constant.ErrorPrefix); found {
			if code == constant.UnknownBattery {
				return nil, not_found.New(constant.BatteryKind, subject)
			}

			if code == constant.StaleReadings || code == constant.DriverNotConnected {
				return nil, detached.New(subject, code)
			}

			return nil, fmt.Errorf(constant.RefusedFormat, command, code)
		}

		if strings.HasPrefix(line, constant.EndPrefix) {
			return result, nil
		}

		result = append(result, line)
	}

	if e = s.Err(); e != nil {
		return nil, e
	}

	return nil, io.ErrUnexpectedEOF
}
