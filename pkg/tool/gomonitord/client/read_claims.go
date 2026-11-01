package client

import (
	"bufio"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/dropped"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
	"io"
	"strings"
)

func ReadClaims(
	r io.Reader,
	receive func([]client.Claim),
) error {
	s := bufio.NewScanner(r)

	for s.Scan() {
		payload, found := strings.CutPrefix(s.Text(), constant.EventPrefix)

		if !found {
			continue
		}

		var v []client.Claim

		if f := notation.Decode(payload, &v); f != nil {
			return fmt.Errorf("stream event: %w", f)
		}

		receive(v)
	}

	if e := s.Err(); e != nil {
		return e
	}

	return dropped.Format("%s", constant.StreamEnded)
}
