package prometheus_server

import "net/http/httptest"

type Server struct {
	*httptest.Server
	status int
	body   string
}
