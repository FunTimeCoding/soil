package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/score_result"
)

func ScoreStatus(v *score_result.Result) *server.ScoreStatus {
	result := &server.ScoreStatus{
		Filled:   v.Filled,
		Mirrored: v.Mirrored,
		Start:    v.Start,
		Total:    v.Total,
	}

	if !v.Synced.IsZero() {
		result.Synced = &v.Synced
	}

	return result
}
