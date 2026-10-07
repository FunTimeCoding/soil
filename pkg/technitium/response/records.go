package response

import "github.com/funtimecoding/soil/pkg/technitium/record"

type Records struct {
	Records []*record.Record `json:"records"`
}
