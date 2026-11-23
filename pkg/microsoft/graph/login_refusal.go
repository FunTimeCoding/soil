package graph

import (
	"context"
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/microsoft/constant"
	"golang.org/x/oauth2"
)

func loginRefusal(e error) error {
	if errors.Is(e, context.DeadlineExceeded) {
		return validation.New(
			"%s: %s",
			constant.SignInUnfinished,
			constant.ExpiredDeviceCode,
		)
	}

	r, okay := errors.AsType[*oauth2.RetrieveError](e)

	if !okay {
		return fmt.Errorf("wait for sign-in: %w", e)
	}

	switch r.ErrorCode {
	case constant.AccessDenied,
		constant.ExpiredDeviceCode,
		constant.DeclinedAuthorization:

		return validation.New(
			"%s: %s",
			constant.SignInUnfinished,
			refusalCode(r),
		)
	default:
		return fmt.Errorf("wait for sign-in: %w", e)
	}
}
