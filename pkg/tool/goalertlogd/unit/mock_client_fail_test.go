package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"testing"
	"time"
)

func TestTheMockFailsOneMethodByItsValue(t *testing.T) {
	m := mock_client.New()
	defer m.Close()
	m.SeedAlert("DiskFull", nil, time.Now())
	m.Fail(m.GetAlertsAfter, unexpected.Format("page refused"))
	_, e := m.GetAlertsAfter(t.Context(), &client.GetAlertsAfterParams{})
	assert.StringContains(t, "page refused", e.Error())
	status, f := m.GetStatus(t.Context())
	assert.FatalOnError(t, f)
	assert.Integer(t, 1, status.LatestIdentifier)
	alerts, g := m.GetAlertsByIdentifier(
		t.Context(),
		&client.GetAlertsByIdentifierParams{Identifier: []int{1}},
	)
	assert.FatalOnError(t, g)
	assert.Count(t, 1, alerts)
}
