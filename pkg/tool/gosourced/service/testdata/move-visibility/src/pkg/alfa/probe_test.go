package alfa

import "testing"

func TestProbe(t *testing.T) {
	if probe() != 4 {
		t.Fail()
	}
}
