package unit

import (
	"example/alfa"
	"testing"
)

func TestOnly(t *testing.T) {
	_ = alfa.Server{Port: 1}
}
