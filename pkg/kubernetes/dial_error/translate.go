package dial_error

import (
	"github.com/funtimecoding/soil/pkg/errors/connection"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
	"strings"
)

func Translate(
	e error,
	host string,
) error {
	if e == nil || connection.Classify(e) != nil {
		return e
	}

	message := e.Error()

	switch {
	case strings.Contains(message, constant.DialRefusedText):
		return connection.New(
			errorsConstant.Unreachable,
			host,
			"",
			errorsConstant.Refused,
		)
	case strings.Contains(message, constant.DialUnknownHostText):
		return connection.New(
			errorsConstant.Unreachable,
			host,
			"",
			errorsConstant.UnknownHost,
		)
	case strings.Contains(message, constant.DialTimeoutText):
		return connection.New(
			errorsConstant.Timeout,
			host,
			"",
			errorsConstant.TimedOut,
		)
	}

	return e
}
