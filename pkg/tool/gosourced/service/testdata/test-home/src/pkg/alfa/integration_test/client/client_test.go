package client

import (
	"example/pkg/alfa"
	"testing"
)

func TestClient(t *testing.T) {
	if alfa.Parse("client") != "client" {
		t.Fatal("client")
	}
}
