package flow

import "testing"

func TestFlowStart(t *testing.T) {
	if start() != 1 {
		t.Fail()
	}
}

func start() int {
	return 1
}
