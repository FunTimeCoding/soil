package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
)

func validateClaim(
	item string,
	owner string,
) error {
	if item == "" || owner == "" {
		return validation.New(constant.ClaimRequiredText)
	}

	return nil
}
