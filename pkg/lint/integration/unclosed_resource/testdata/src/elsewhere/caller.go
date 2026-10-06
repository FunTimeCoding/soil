package example

import (
	"example/helper"
	"testing"
)

func UseArrangedElsewhere(t *testing.T) {
	s := helper.NewServer(t)
	s.Ping()
}
