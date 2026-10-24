package requester

import (
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"time"
)

func (r *Requester) wait(q *request.Request) error {
	if r.pace == nil {
		return nil
	}

	reservation := r.pace.Reserve()
	timer := time.NewTimer(reservation.Delay())
	defer timer.Stop()

	select {
	case <-q.Context.Done():
		reservation.Cancel()

		return q.Context.Err()
	case <-timer.C:
		return nil
	}
}
