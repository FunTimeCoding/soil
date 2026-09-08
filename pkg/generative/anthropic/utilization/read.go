package utilization

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"io"
	"net/http"
)

func Read(token string) *Result {
	r, e := http.NewRequest(
		http.MethodGet,
		constant.AnthropicUtilizationLink,
		nil,
	)
	errors.PanicOnError(e)
	r.Header.Set(web.Authorization, key_value.Space(web.Bearer, token))
	r.Header.Set(
		constant.AnthropicBetaHeader,
		constant.AnthropicUtilizationBeta,
	)
	response, f := http.DefaultClient.Do(r)
	errors.PanicOnError(f)
	defer errors.PanicClose(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil
	}

	body, g := io.ReadAll(response.Body)
	errors.PanicOnError(g)

	return Parse(body)
}
