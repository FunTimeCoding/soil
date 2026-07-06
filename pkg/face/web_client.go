package face

import "github.com/funtimecoding/soil/pkg/web/web_client/web_response"

type WebClient interface {
	IsWebClient()
	Get(locator string) (*web_response.Response, error)
	Post(
		locator string,
		body any,
	) (*web_response.Response, error)
}
