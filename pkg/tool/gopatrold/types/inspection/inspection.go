package inspection

import (
	"github.com/funtimecoding/soil/pkg/mdstat/array"
	"github.com/funtimecoding/soil/pkg/smartctl/response"
)

type Inspection struct {
	Host    string
	Reached bool
	Arrays  []*array.Array
	Drives  []*response.Report
}
