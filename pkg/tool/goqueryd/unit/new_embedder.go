package unit

import "testing"

func newEmbedder(t *testing.T) *FixedEmbedder {
	t.Helper()
	result := new(FixedEmbedder)
	t.Cleanup(result.Close)

	return result
}
