package technitium

import "github.com/funtimecoding/soil/pkg/technitium/record"

type addRecordResponse struct {
	AddedRecord *record.Record `json:"addedRecord"`
}
