package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization"
	"testing"
)

func TestUtilizationParse(t *testing.T) {
	result := utilization.Parse(
		[]byte(fixture.Read("claude", "utilization.json")),
	)
	assert.NotNil(t, result)
	assert.Integer(t, 4, result.SessionPercent)
	assert.String(
		t,
		"2026-09-08T18:49:59Z",
		result.SessionReset.UTC().Format("2006-01-02T15:04:05Z"),
	)
	assert.Integer(t, 22, result.WeeklyPercent)
	assert.String(
		t,
		"2026-09-15T14:59:59Z",
		result.WeeklyReset.UTC().Format("2006-01-02T15:04:05Z"),
	)
	assert.Integer(t, 34, result.FablePercent)
	assert.True(t, result.FableSeen)
	assert.False(t, result.FableReset.IsZero())
}

func TestUtilizationParseUnusedFable(t *testing.T) {
	result := utilization.Parse(
		[]byte(fixture.Read("claude", "utilization-unused-fable.json")),
	)
	assert.NotNil(t, result)
	assert.Integer(t, 3, result.SessionPercent)
	assert.Integer(t, 0, result.FablePercent)
	assert.True(t, result.FableSeen)
	assert.True(t, result.FableReset.IsZero())
}

func TestUtilizationParseEmpty(t *testing.T) {
	assert.Nil(t, utilization.Parse([]byte("{}")))
	assert.Nil(t, utilization.Parse([]byte("not json")))
}

func TestUtilizationParseCredential(t *testing.T) {
	result := utilization.ParseCredential(
		fixture.Read("claude", "credentials.json"),
	)
	assert.NotNil(t, result)
	assert.String(t, "sk-ant-access-placeholder", result.AccessToken)
	assert.String(t, "team", result.SubscriptionType)
	assert.String(t, "default_raven", result.RateLimitTier)
	assert.String(
		t,
		"2026-09-08T22:41:34Z",
		result.ExpiresAt.UTC().Format("2006-01-02T15:04:05Z"),
	)
	assert.True(t, result.Expired(result.ExpiresAt.Add(1)))
	assert.False(t, result.Expired(result.ExpiresAt.Add(-1)))
}

func TestUtilizationParseCredentialEmpty(t *testing.T) {
	assert.Nil(t, utilization.ParseCredential(""))
	assert.Nil(t, utilization.ParseCredential("not json"))
	assert.Nil(t, utilization.ParseCredential(`{"claudeAiOauth":{}}`))
}
