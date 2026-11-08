package unit

import (
	"example/pkg/service"
	"testing"
)

func TestPlace(t *testing.T) {
	var s service.Service
	s.Place("charlie")
}
