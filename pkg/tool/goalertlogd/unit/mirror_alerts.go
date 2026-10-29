package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"testing"
	"time"
)

func mirrorAlerts(
	t *testing.T,
	c face.Client,
) *mock_client.Client {
	t.Helper()
	result := mock_client.New()
	limit := constant.PageMaximum
	cursor := 0

	for {
		page, e := c.GetAlertsAfter(
			t.Context(),
			&client.GetAlertsAfterParams{Identifier: cursor, Limit: &limit},
		)
		assert.FatalOnError(t, e)

		for _, a := range page {
			start, f := time.Parse(constant.DateFormat, a.Start)
			assert.FatalOnError(t, f)
			result.SeedAlert(a.Name, a.Labels, start)

			if a.End != nil {
				end, g := time.Parse(constant.DateFormat, *a.End)
				assert.FatalOnError(t, g)
				result.Resolve(a.Identifier, end)
			}

			cursor = a.Identifier
		}

		if len(page) < limit {
			return result
		}
	}
}
