package response

import "github.com/funtimecoding/soil/pkg/technitium/record"

type AddRecord struct {
	AddedRecord *record.Record `json:"addedRecord"`
}
