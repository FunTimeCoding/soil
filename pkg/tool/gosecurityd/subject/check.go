package subject

import "github.com/funtimecoding/soil/pkg/errors/validation"

func Check(s string) error {
	if Kind(s) != "" {
		return nil
	}

	return validation.New(
		"subject malformed, expect a CVE identifier, upgrade:<release>:<source>, release:<host>, reboot:<host> or kernel:<host>: %s",
		s,
	)
}
