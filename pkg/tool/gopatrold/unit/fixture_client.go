package unit

import (
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/mdstat"
	"github.com/funtimecoding/soil/pkg/patrol/mock_client"
	"github.com/funtimecoding/soil/pkg/smartctl"
)

func fixtureClient(
	mdstatName string,
	reportNames ...string,
) *mock_client.Client {
	result := mock_client.New()
	result.AddArray(mdstat.Parse(fixture.Read("disk", mdstatName))...)

	for _, name := range reportNames {
		r, e := smartctl.Report(fixture.Read("disk", name))
		errors.PanicOnError(e)
		result.AddDrive(r)
	}

	return result
}
