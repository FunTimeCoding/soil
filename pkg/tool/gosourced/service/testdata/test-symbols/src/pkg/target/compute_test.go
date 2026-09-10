package target

import "testing"

func TestCompute(t *testing.T) {
	if Compute() != 1 {
		t.Fail()
	}
}
