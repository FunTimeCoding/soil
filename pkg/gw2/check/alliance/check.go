package alliance

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/system"
)

func Check() {
	a := argument.NewSimple("alliance")
	a.String(constant.Tag, "", "Guild tag")
	a.ParseSimple()
	Log(
		fmt.Sprintf("%s\\AppData\\Local\\ArcdpsLogManager", system.Home()),
		a.GetString(constant.Tag),
	)
}
