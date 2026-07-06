package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestTimelinePagination(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()

	for i := 1; i <= 12; i++ {
		a.Announce(a.Name(), fmt.Sprintf("topic-%02d", i))
	}

	x := context.Background()
	first, e := a.RestClient.GetTimelineWithResponse(
		x,
		&client.GetTimelineParams{
			Limit:  new(5),
			Offset: new(0),
		},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 5, *first.JSON200)
	assert.StringContains(t, "topic-12", (*first.JSON200)[0].Subject)
	assert.StringContains(t, "topic-08", (*first.JSON200)[4].Subject)
	second, f := a.RestClient.GetTimelineWithResponse(
		x,
		&client.GetTimelineParams{
			Limit:  new(5),
			Offset: new(5),
		},
	)
	assert.FatalOnError(t, f)
	assert.Count(t, 5, *second.JSON200)
	assert.StringContains(t, "topic-07", (*second.JSON200)[0].Subject)
	assert.StringContains(t, "topic-03", (*second.JSON200)[4].Subject)
	third, g := a.RestClient.GetTimelineWithResponse(
		x,
		&client.GetTimelineParams{
			Limit:  new(5),
			Offset: new(10),
		},
	)
	assert.FatalOnError(t, g)
	assert.Count(t, 2, *third.JSON200)
	assert.StringContains(t, "topic-02", (*third.JSON200)[0].Subject)
	assert.StringContains(t, "topic-01", (*third.JSON200)[1].Subject)
}
