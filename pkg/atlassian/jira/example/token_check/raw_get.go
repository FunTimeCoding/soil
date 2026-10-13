package token_check

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"io"
	"net/http"
)

func rawGet(
	locator string,
	user string,
	token string,
) {
	r, e := http.NewRequest(http.MethodGet, locator, nil)
	errors.PanicOnError(e)
	r.SetBasicAuth(user, token)
	p := stall_client.New().Build()
	defer p.Close()
	response, f := p.Client().Do(r)
	errors.PanicOnError(f)
	defer errors.PanicClose(response.Body)
	body, g := io.ReadAll(response.Body)
	errors.PanicOnError(g)
	console.Format("  Status: %d\n", response.StatusCode)
	console.Format("  Body (first 500): %.500s\n", body)
}
