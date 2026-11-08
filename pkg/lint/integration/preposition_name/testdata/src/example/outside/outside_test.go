package outside

import "testing"

func TestItemByNameForEmptyStore(t *testing.T) {
	if new(Store).ItemByName("x") != nil {
		t.Fatal("found")
	}
}
