package bravo_test

import (
	"example/alfa"
	"testing"
)

func TestExternal(t *testing.T) {
	alfa.NewServer().Start()
}
