package sublime

import (
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
)

type Client struct {
	base   *locator.Locator
	client *http.Client
}
