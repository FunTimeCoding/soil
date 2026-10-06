package connection

import "errors"

func Classify(e error) *Failure {
	if e == nil {
		return nil
	}

	var existing *Failure

	if errors.As(e, &existing) {
		return existing
	}

	k, reason := kind(e)

	if k == "" {
		return nil
	}

	host, path := locate(e)

	return New(k, host, path, reason)
}
