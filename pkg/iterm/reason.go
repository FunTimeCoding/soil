package iterm

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func reason(r *http.Response) string {
	return refusal.Reason(system.ReadAll(r.Body), r.Status)
}
