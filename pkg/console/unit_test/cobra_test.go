package unit_test

import (
	"github.com/funtimecoding/soil/pkg/console/cobra"
	"testing"
)

func TestCobra(t *testing.T) {
	cobra.Execute(cobra.New("ls"))
}
