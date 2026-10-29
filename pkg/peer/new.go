package peer

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer/constant"
	"github.com/funtimecoding/soil/pkg/peer/reacher"
	"time"
)

func New(
	name string,
	address string,
	p face.Reachable,
	l *logger.Logger,
	r face.Reporter,
) *Guard {
	x, cancel := context.WithCancel(context.Background())

	return &Guard{
		name:     name,
		address:  address,
		probe:    p,
		reacher:  reacher.New(r),
		logger:   l,
		reporter: r,
		minimum:  constant.ProbeMinimum,
		maximum:  constant.ProbeMaximum,
		wake:     make(chan struct{}, 1),
		context:  x,
		cancel:   cancel,
		clock:    time.Now,
	}
}
