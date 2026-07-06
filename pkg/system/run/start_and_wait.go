package run

import "os/exec"

func (r *Run) startAndWait(c *exec.Cmd) error {
	if r.registry == nil {
		return c.Run()
	}

	e := c.Start()

	if e != nil {
		return e
	}

	r.registry.Add(c.Process.Pid)
	e = c.Wait()
	r.registry.Remove(c.Process.Pid)

	return e
}
