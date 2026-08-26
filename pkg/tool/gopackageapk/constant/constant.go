package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"gopackageapk",
	"Alpine APK package builder",
	"gopackageapk <executable> <version>",
)
