package unit

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
	"testing"
)

func TestMockAnswersStatusItWasGiven(t *testing.T) {
	m := mock_client.New()
	defer m.Close()
	m.SetStatus(&client.Status{Instance: "alfa", Power: "on"})
	s, e := m.GetStatus(
		context.Background(),
		&client.GetStatusParams{Instance: new("alfa")},
	)
	assert.FatalOnError(t, e)
	assert.String(t, "on", s.Power)
	_, f := m.GetStatus(
		context.Background(),
		&client.GetStatusParams{Instance: new("bravo")},
	)
	assert.True(t, not_found.Is(f))
}

func TestMockRecordsPowerRequestsUntilTold(t *testing.T) {
	m := mock_client.New()
	defer m.Close()
	m.SetStatus(&client.Status{Instance: "alfa", Power: "on"})
	r, e := m.RequestPower(
		context.Background(),
		client.PowerRequest{Instance: new("alfa"), Operation: "shutdown"},
	)
	assert.FatalOnError(t, e)
	assert.String(t, "alfa", r.Instance)
	assert.String(t, "off soft graceful", r.Requested)
	assert.String(t, "on", r.Power)
	_, n := m.RequestPower(
		context.Background(),
		client.PowerRequest{Instance: new("bravo"), Operation: "start"},
	)
	assert.True(t, not_found.Is(n))
	m.Fail(m.RequestPower, errors.New("engine busy"))
	_, f := m.RequestPower(
		context.Background(),
		client.PowerRequest{Instance: new("alfa"), Operation: "start"},
	)
	assert.NotNil(t, f)
	assert.Count(t, 1, m.Requested)
}
