package example

import "os/exec"

func SuppressedBesideSecurity(s []string) {
	exec.Command(s[0], s[1:]...) // #nosec G204 goanalyze:ignore forbidden_call
}
