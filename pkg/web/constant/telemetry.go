package constant

const (
	// https://opentelemetry.io/docs/specs/semconv/http/http-spans/
	// Required attributes
	TelemetryRequestMethod = "http.request.method"
	TelemetryPath          = "url.path"
	TelemetryScheme        = "url.scheme"
	// Conditionally required
	TelemetryQuery  = "url.query"
	TelemetryRoute  = "http.route"
	TelemetryStatus = "http.response.status_code"
	// Recommended
	TelemetryClient    = "client.address"
	TelemetryPeer      = "network.peer.address"
	TelemetryProtocol  = "network.protocol.version"
	TelemetryServer    = "server.address"
	TelemetryUserAgent = "user_agent.original"
	// Opt-in attributes for webhooks
	TelemetryBodySize     = "http.request.body.size"
	TelemetryBody         = "http.request.body"
	TelemetryHeaderPrefix = "http.request.header"
)
