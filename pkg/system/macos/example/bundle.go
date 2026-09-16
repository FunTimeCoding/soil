package example

import (
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/macos"
)

func Bundle() {
	macos.CreateBundle(
		"Example",
		constant.Temporary,
		"tmp/example",
		"tmp/icon.icns",
		"com.example",
		library.DefaultVersion,
	)
}
