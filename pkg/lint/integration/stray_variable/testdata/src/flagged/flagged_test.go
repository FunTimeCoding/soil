package flagged

import "testing"

var testFixture = "fixture"

func TestFlagged(t *testing.T) {
	if testFixture == "" {
		t.Fail()
	}
}
