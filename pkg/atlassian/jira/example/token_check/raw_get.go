package token_check

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
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
	response, f := http.DefaultClient.Do(r)
	errors.PanicOnError(f)
	defer errors.PanicClose(response.Body)
	body, g := io.ReadAll(response.Body)
	errors.PanicOnError(g)
	fmt.Printf("  Status: %d\n", response.StatusCode)
	fmt.Printf("  Body (first 500): %.500s\n", body)
}
