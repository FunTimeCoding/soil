package example

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system/macos"
)

func Notify() {
	macos.Notify(
		stringsConstant.UpperAlfa,
		stringsConstant.UpperBravo,
		stringsConstant.UpperCharlie,
	)
	macos.SimpleNotify(stringsConstant.UpperAlfa)
	macos.Beep()
	macos.Alert("Subject", "Body")
	macos.InputDialog("Test1", "Test2", "")
	macos.CustomDialog("Test1", "Test2")
}
