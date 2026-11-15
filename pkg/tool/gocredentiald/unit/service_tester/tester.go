package service_tester

import (
	"github.com/funtimecoding/soil/pkg/face"
	"time"
)

type Tester struct {
	revealed []string
	now      func() time.Time
	reporter face.Reporter
}
