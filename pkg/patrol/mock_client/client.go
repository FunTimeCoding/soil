package mock_client

import (
	"github.com/funtimecoding/soil/pkg/mdstat/array"
	"github.com/funtimecoding/soil/pkg/smartctl/response"
)

type Client struct {
	Failure error
	Calls   int
	arrays  []*array.Array
	drives  []*response.Report
}
