package result

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"

func NewFailed(r *client.FailedReport) *Failed {
	return &Failed{Report: r, Summary: &FailedSummary{Refused: len(r.Chunks)}}
}
