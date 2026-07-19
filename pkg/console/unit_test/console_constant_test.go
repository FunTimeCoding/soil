package unit_test

import (
	"github.com/funtimecoding/soil/pkg/console"
	"testing"
)

func TestConsoleConstant(t *testing.T) {
	console.Blue("%s", "test")
	console.Cyan("%s", "test")
	console.Green("%s", "test")
	console.Magenta("%s", "test")
	console.Red("%s", "test")
	console.Yellow("%s", "test")
}
