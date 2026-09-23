package plain

import "testing"

type Holder struct{}

func (o *Holder) Take(
	t *testing.T,
	name string,
) string {
	t.Helper()

	return name
}
