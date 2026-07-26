package example

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system/macos"
)

func Notify() {
	macos.Notify(constant.UpperAlfa, constant.UpperBravo, constant.UpperCharlie)
	macos.SimpleNotify(constant.UpperAlfa)
	macos.Beep()
	macos.Alert("Subject", "Body")
	macos.InputDialog("Test1", "Test2", "")
	macos.CustomDialog("Test1", "Test2")
}
