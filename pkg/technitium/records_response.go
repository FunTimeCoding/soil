package technitium

import "github.com/funtimecoding/soil/pkg/technitium/record"

type recordsResponse struct {
	Records []*record.Record `json:"records"`
}
