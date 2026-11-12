package constant

const (
	PortEnvironment = "METRIC_PORT"
	PortUsage       = "Metric listen port"
	BindEnvironment = "METRIC_BIND_ADDRESS"
	BindUsage       = "Metric listen address, every interface when empty"

	Port = 9090

	UpstreamMetric = "upstream_request_duration_seconds"
	UpstreamHelp   = "Calls to upstream hosts by response code, or failure kind when no response arrived"
	IdentityLabel  = "identity"
)
