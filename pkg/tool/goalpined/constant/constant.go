package constant

import "github.com/funtimecoding/soil/pkg/identity"

var Identity = identity.New(
	"goalpined",
	"Alpine package repository server",
	"goalpined",
).WithInstructions(
	"Alpine package repository server - upload, sign, and index apk packages. Use list_packages to see what the repository index holds.",
)

const (
	ListPackages = "list_packages"

	FileAddress = ":8081"
)

const (
	HostEnvironment         = "GOALPINE_HOST"
	PortEnvironment         = "GOALPINE_PORT"
	InsecureEnvironment     = "GOALPINE_INSECURE"
	TokenEnvironment        = "GOALPINE_TOKEN"
	SignatureKeyEnvironment = "ALPINE_SIGNATURE_KEY"

	FixtureUploadPath = "/api/packages/v3.21/main/x86_64/a.apk"

	PackageRoot  = "/apk"
	KeyDirectory = "/key"
)
