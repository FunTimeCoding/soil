package assert

import (
	"net/http"
	"testing"
)

func HTTPStatus(
	t *testing.T,
	c *http.Client,
	locator string,
	expect int,
) {
	t.Helper()
	r, e := c.Get(locator)

	if e != nil {
		t.Errorf("GET %s: %s", locator, e)

		return
	}

	defer func() {
		if f := r.Body.Close(); f != nil {
			panic(f)
		}
	}()
	Integer(t, expect, r.StatusCode)
}
