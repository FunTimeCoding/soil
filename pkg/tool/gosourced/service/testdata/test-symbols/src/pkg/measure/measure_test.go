package measure

import "testing"

func TestMeasure(t *testing.T) {
	if Measure() != 2 {
		t.Fail()
	}
}
