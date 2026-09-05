package run

import (
	"os"
	"os/exec"
)

func (r *Run) build(s ...string) *exec.Cmd {
	c := exec.Command(s[0], s[1:]...) // goanalyze:ignore

	if r.Directory != "" {
		c.Dir = r.Directory
	}

	if r.Input != nil {
		c.Stdin = r.Input
	}

	if len(r.environment) > 0 {
		if r.replaceEnviron {
			c.Env = r.environment
		} else {
			c.Env = append(os.Environ(), r.environment...)
		}
	}

	if r.processGroup {
		setProcessGroup(c)
	}

	if r.stdout != nil {
		c.Stdout = r.stdout
	}

	if r.stderr != nil {
		c.Stderr = r.stderr
	}

	return c
}
