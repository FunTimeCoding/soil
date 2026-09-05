package machinery

import "testing"

// golint:fixture stray_constant
const testFixture = "fixture"

func TestMachinery(t *testing.T) {
	if testFixture == "" {
		t.Fail()
	}
}
