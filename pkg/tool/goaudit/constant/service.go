package constant

const (
	ServiceFlag            = "service"
	MissingServiceKey      = "missing_service"
	MissingServiceText     = "lifecycle daemon without service/ - see service-tool.md, Service"
	RunBuildsClientKey     = "run_builds_client"
	RunBuildsClientText    = "run.go builds %s - the service builds it from a factory on the option"
	ServiceNotClosedKey    = "service_not_closed"
	ServiceNotClosedText   = "run.go builds the service but never defers its Close"
	ServiceDirectory       = "service"
	ServiceCloseFile       = "close.go"
	LifecyclePackage       = "lifecycle"
	ServicePackage         = "service"
	CloseMethod            = "Close"
	OpenPrefix             = "Open"
)

var RunPlumbing = []string{
	"guard",
	"lifecycle",
	"logger",
	"metric",
	"option",
	"server",
	"service",
	"ticker",
	"web",
	"worker",
}
