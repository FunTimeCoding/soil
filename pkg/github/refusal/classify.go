package refusal

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"github.com/google/go-github/v93/github"
)

func Classify(e error) error {
	if e == nil {
		return nil
	}

	if f := connection.Classify(e); f != nil {
		return f
	}

	var rate *github.RateLimitError
	var abuse *github.AbuseRateLimitError

	if errors.As(e, &rate) || errors.As(e, &abuse) {
		return unavailable.Format("%s", e.Error())
	}

	var answer *github.ErrorResponse

	if !errors.As(e, &answer) || answer.Response == nil {
		return e
	}

	if g := refusal.Status(answer.Response.StatusCode, answer.Message); g != nil {
		return g
	}

	return e
}
