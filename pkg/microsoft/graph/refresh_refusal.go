package graph

import (
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/funtimecoding/soil/pkg/microsoft/constant"
	"golang.org/x/oauth2"
)

func refreshRefusal(e error) error {
	r, okay := errors.AsType[*oauth2.RetrieveError](e)

	if !okay {
		return fmt.Errorf("refresh graph token: %w", e)
	}

	switch r.ErrorCode {
	case constant.ExpiredGrant, constant.InteractionRequired:
		return not_configured.Format(
			"%s: %s",
			constant.SignInExpired,
			refusalCode(r),
		)
	case constant.InvalidClient, constant.UnauthorizedClient:
		return not_configured.Format(
			"%s: %s",
			constant.RegistrationRefused,
			refusalCode(r),
		)
	default:
		return fmt.Errorf("refresh graph token: %w", e)
	}
}
