package connection

import (
	"github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/dropped"
	"github.com/funtimecoding/soil/pkg/errors/timeout"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
)

func New(
	kind string,
	host string,
	path string,
	reason string,
) *Failure {
	result := &Failure{Kind: kind, Host: host, Path: path, Reason: reason}

	switch kind {
	case constant.Timeout:
		result.class = timeout.Format("%s", result.Error())
	case constant.Dropped:
		result.class = dropped.Format("%s", result.Error())
	default:
		result.class = unreachable.Format("%s", result.Error())
	}

	return result
}
