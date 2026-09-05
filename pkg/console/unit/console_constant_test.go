package unit

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"testing"
)

func TestConsoleConstant(t *testing.T) {
	constant.Blue("%s", "test")
	constant.Cyan("%s", "test")
	constant.Green("%s", "test")
	constant.Magenta("%s", "test")
	constant.Red("%s", "test")
	constant.Yellow("%s", "test")
}
