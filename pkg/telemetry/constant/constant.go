package constant

import "time"

const CommandLineTimeout = 500 * time.Millisecond
const (
	ModelContext = "model_context"
	CommandLine  = "command_line"
	Web          = "web"
	WebService   = "web_service"

	User = "user"

	Success = "success"
	Error   = "error"
	Blocked = "blocked"

	Baseline = "baseline"
	Domain   = "domain"

	HostEnvironment     = "GOTELEMETRY_HOST"
	PortEnvironment     = "GOTELEMETRY_PORT"
	InsecureEnvironment = "GOTELEMETRY_INSECURE"
)
