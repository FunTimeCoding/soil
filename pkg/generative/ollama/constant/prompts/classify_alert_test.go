package prompts

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestClassifyAlert(t *testing.T) {
	// Classify A or B
	assert.String(
		t,
		`Instructions: Decide if this Prometheus alert is already-broken or not-yet-broken
Sample to classify: ProbeFail

## Answer format
JSON, 2 string keys: Reason, Answer
Reason: One concise reasoning sentence.
Answer: Choice between already-broken, not-yet-broken

## Example data
already-broken: DiskFull
not-yet-broken: DiskNearFull,LatencyHigh`,
		ClassifyAlert().Render(),
	)
}
