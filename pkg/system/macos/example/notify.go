package example

import (
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/funtimecoding/soil/pkg/system/macos"
)

func Notify() {
	macos.Notify(upper.Alfa, upper.Bravo, upper.Charlie)
	macos.SimpleNotify(upper.Alfa)
	macos.Beep()
	macos.Alert("Subject", "Body")
	macos.InputDialog("Test1", "Test2", "")
	macos.CustomDialog("Test1", "Test2")
}
