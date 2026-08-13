package embed

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/generative/openai"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/url"
)

func NewEnvironment() face.Embedder {
	target := environment.Required(constant.EmbedTargetEnvironment)
	u, e := url.Parse(target)
	errors.PanicOnError(e)

	switch u.Scheme {
	case constant.OllamaScheme:
		options := []ollama.Option{ollama.WithHost(u.Hostname())}

		if u.Port() != "" {
			options = append(
				options,
				ollama.WithPort(strings.MustToInteger(u.Port())),
			)
		}

		return ollama.New(options...)
	case constant.OpenAIScheme:
		l := locator.New(u.Hostname()).Insecure().Path("/v1")

		if u.Port() != "" {
			l.Port(strings.MustToInteger(u.Port()))
		}

		return openai.NewBase(l.String())
	}

	panic(fmt.Sprintf("unknown embed target scheme: %s", u.Scheme))
}
