package client

import (
	"testing"
	"time"
)

func streamEnd(
	t *testing.T,
	ended chan error,
) error {
	t.Helper()

	select {
	case e := <-ended:
		return e
	case <-time.After(5 * time.Second):
		t.Fatal("stream did not end within 5 seconds")

		return nil
	}
}
