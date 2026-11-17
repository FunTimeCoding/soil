package base

import (
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/mdstat"
	"github.com/funtimecoding/soil/pkg/smartctl"
)

func (s *Server) SeedDisks(
	instance string,
	mdstatName string,
	reportNames ...string,
) {
	s.t.Helper()
	c := s.Instance(instance)
	c.AddArray(mdstat.Parse(fixture.Read("disk", mdstatName))...)

	for _, name := range reportNames {
		r, e := smartctl.Report(fixture.Read("disk", name))
		errors.PanicOnError(e)
		c.AddDrive(r)
	}
}
