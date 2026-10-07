package log

import (
	"github.com/funtimecoding/soil/pkg/gw2/log_manager/response"
	"time"
)

type Log struct {
	Accounts []string
	Time     time.Time
	Raw      *response.Log
}
