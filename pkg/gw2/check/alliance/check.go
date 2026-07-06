package alliance

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/system"
)

func Check() {
	path := fmt.Sprintf(
		"%s\\AppData\\Local\\ArcdpsLogManager",
		system.Home(),
	)

	if false {
		Guild(path)
	}

	if true {
		a := argument.NewSimple("alliance")
		a.String(argument.Tag, "", "Guild tag")
		a.ParseSimple()
		Log(path, a.GetString(argument.Tag))
	}
}
