package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/check/silence/matcher"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/prometheus/unit/silence_tester"
	"github.com/prometheus/alertmanager/api/v2/models"
	"testing"
	"time"
)

func TestMatchesExactExistingLabel(t *testing.T) {
	assert.True(
		t,
		matchesSingle(
			"job",
			"test",
			true,
			false,
			models.LabelSet{"job": "test", "alertname": "HighCPU"},
		),
	)
}

func TestMatchesExactDifferentValue(t *testing.T) {
	assert.False(
		t,
		matchesSingle(
			"job",
			"test",
			true,
			false,
			models.LabelSet{"job": "production", "alertname": "HighCPU"},
		),
	)
}

func TestMatchesExactMissingLabel(t *testing.T) {
	assert.False(
		t,
		matchesSingle(
			"job",
			"test",
			true,
			false,
			models.LabelSet{"alertname": "HighCPU"},
		),
	)
}

func TestMatchesExactEmptyMissingLabel(t *testing.T) {
	assert.True(
		t,
		matchesSingle(
			"job",
			"",
			true,
			false,
			models.LabelSet{"alertname": "HighCPU"},
		),
	)
}

func TestMatchesExactEmptyExistingEmptyLabel(t *testing.T) {
	assert.True(
		t,
		matchesSingle(
			"job",
			"",
			true,
			false,
			models.LabelSet{"job": "", "alertname": "HighCPU"},
		),
	)
}

func TestMatchesNotEqualMatchingValue(t *testing.T) {
	assert.False(
		t,
		matchesSingle(
			"job",
			"test",
			false,
			false,
			models.LabelSet{"job": "test", "alertname": "HighCPU"},
		),
	)
}

func TestMatchesNotEqualDifferentValue(t *testing.T) {
	assert.True(
		t,
		matchesSingle(
			"job",
			"test",
			false,
			false,
			models.LabelSet{"job": "production", "alertname": "HighCPU"},
		),
	)
}

func TestMatchesNotEqualMissingLabel(t *testing.T) {
	assert.True(
		t,
		matchesSingle(
			"job",
			"test",
			false,
			false,
			models.LabelSet{"alertname": "HighCPU"},
		),
	)
}

func TestMatchesNotEqualEmptyMissingLabel(t *testing.T) {
	assert.False(
		t,
		matchesSingle(
			"job",
			"",
			false,
			false,
			models.LabelSet{"alertname": "HighCPU"},
		),
	)
}

func TestMatchesNotEqualEmptyNonEmptyLabel(t *testing.T) {
	assert.True(
		t,
		matchesSingle(
			"job",
			"",
			false,
			false,
			models.LabelSet{"job": "test", "alertname": "HighCPU"},
		),
	)
}

func TestMatchesRegexMatchingValue(t *testing.T) {
	assert.True(
		t,
		matchesSingle(
			"job",
			"test.*",
			true,
			true,
			models.LabelSet{"job": "test-job", "alertname": "HighCPU"},
		),
	)
}

func TestMatchesRegexNonMatchingValue(t *testing.T) {
	assert.False(
		t,
		matchesSingle(
			"job",
			"test.*",
			true,
			true,
			models.LabelSet{"job": "production", "alertname": "HighCPU"},
		),
	)
}

func TestMatchesRegexMissingLabel(t *testing.T) {
	assert.False(
		t,
		matchesSingle(
			"job",
			"test.*",
			true,
			true,
			models.LabelSet{"alertname": "HighCPU"},
		),
	)
}

func TestMatchesRegexEmptyMissingLabel(t *testing.T) {
	assert.True(
		t,
		matchesSingle(
			"job",
			"",
			true,
			true,
			models.LabelSet{"alertname": "HighCPU"},
		),
	)
}

func TestMatchesRegexWildcardMissingLabel(t *testing.T) {
	assert.True(
		t,
		matchesSingle(
			"job",
			".*",
			true,
			true,
			models.LabelSet{"alertname": "HighCPU"},
		),
	)
}

func TestMatchesRegexWildcardExistingLabel(t *testing.T) {
	assert.True(
		t,
		matchesSingle(
			"job",
			".*",
			true,
			true,
			models.LabelSet{"job": "test", "alertname": "HighCPU"},
		),
	)
}

func TestMatchesNotRegexMatchingValue(t *testing.T) {
	assert.False(
		t,
		matchesSingle(
			"job",
			"test.*",
			false,
			true,
			models.LabelSet{"job": "test-job", "alertname": "HighCPU"},
		),
	)
}

func TestMatchesNotRegexNonMatchingValue(t *testing.T) {
	assert.True(
		t,
		matchesSingle(
			"job",
			"test.*",
			false,
			true,
			models.LabelSet{"job": "production", "alertname": "HighCPU"},
		),
	)
}

func TestMatchesNotRegexMissingLabel(t *testing.T) {
	assert.True(
		t,
		matchesSingle(
			"job",
			"test.*",
			false,
			true,
			models.LabelSet{"alertname": "HighCPU"},
		),
	)
}

func TestMatchesNotRegexWildcardMissingLabel(t *testing.T) {
	assert.False(
		t,
		matchesSingle(
			"job",
			".*",
			false,
			true,
			models.LabelSet{"alertname": "HighCPU"},
		),
	)
}

func TestMatchesSetAllMatch(t *testing.T) {
	assert.True(
		t,
		matchesSet(
			[]*models.Matcher{
				silence_tester.NewMatcher("alertname", "HighCPU", true, false),
				silence_tester.NewMatcher(
					constant.SeverityLabel,
					constant.CriticalSeverity,
					true,
					false,
				),
			},
			models.LabelSet{"alertname": "HighCPU", "severity": "critical"},
		),
	)
}

func TestMatchesSetOneMismatch(t *testing.T) {
	assert.False(
		t,
		matchesSet(
			[]*models.Matcher{
				silence_tester.NewMatcher("alertname", "HighCPU", true, false),
				silence_tester.NewMatcher(
					constant.SeverityLabel,
					constant.WarningSeverity,
					true,
					false,
				),
			},
			models.LabelSet{"alertname": "HighCPU", "severity": "critical"},
		),
	)
}

func TestMatchesSetPositiveAndNegative(t *testing.T) {
	assert.True(
		t,
		matchesSet(
			[]*models.Matcher{
				silence_tester.NewMatcher("alertname", "HighCPU", true, false),
				silence_tester.NewMatcher("job", "test", false, false),
			},
			models.LabelSet{"alertname": "HighCPU", "job": "production"},
		),
	)
}

func TestMatchesSetNegativeMissingLabel(t *testing.T) {
	assert.True(
		t,
		matchesSet(
			[]*models.Matcher{
				silence_tester.NewMatcher("alertname", "HighCPU", true, false),
				silence_tester.NewMatcher("job", "test", false, false),
			},
			models.LabelSet{"alertname": "HighCPU"},
		),
	)
}

func TestMatchesWindowActive(t *testing.T) {
	now := time.Now()
	assert.True(
		t,
		matchesWindow(now.Add(-1*time.Hour), now.Add(1*time.Hour), now),
	)
}

func TestMatchesWindowNotStarted(t *testing.T) {
	now := time.Now()
	assert.False(
		t,
		matchesWindow(now.Add(1*time.Hour), now.Add(2*time.Hour), now),
	)
}

func TestMatchesWindowExpired(t *testing.T) {
	now := time.Now()
	assert.False(
		t,
		matchesWindow(now.Add(-2*time.Hour), now.Add(-1*time.Hour), now),
	)
}

func TestMatchesWindowStartsNow(t *testing.T) {
	now := time.Now()
	assert.True(t, matchesWindow(now, now.Add(1*time.Hour), now))
}

func TestMatchesWindowEndsNow(t *testing.T) {
	now := time.Now()
	assert.False(t, matchesWindow(now.Add(-1*time.Hour), now, now))
}

func TestMatchesWithMultipleAlerts(t *testing.T) {
	now := time.Now()
	s := silence_tester.NewSilence(
		new(now.Add(-1*time.Hour)),
		new(now.Add(1*time.Hour)),
		constant.SeverityLabel,
		constant.CriticalSeverity,
		new(true),
		new(false),
	)
	alerts := []*alert.Alert{
		alert.NewFromLabels(
			models.LabelSet{"alertname": "HighCPU", "severity": "critical"},
		),
		alert.NewFromLabels(
			models.LabelSet{"alertname": "HighMemory", "severity": "warning"},
		),
		alert.NewFromLabels(
			models.LabelSet{"alertname": "DiskFull", "severity": "critical"},
		),
		alert.NewFromLabels(
			models.LabelSet{"alertname": "NetworkIssue", "severity": "info"},
		),
	}
	m := matcher.Matches(s, alerts, now)

	if len(m) != 2 {
		t.Errorf("Expected 2 matching alerts, got %d", len(m))
	}

	for _, a := range m {
		if a.Labels[constant.SeverityLabel] != "critical" {
			t.Errorf(
				"Expected all matched alerts to have severity=critical, got %s",
				a.Labels[constant.SeverityLabel],
			)
		}
	}
}
