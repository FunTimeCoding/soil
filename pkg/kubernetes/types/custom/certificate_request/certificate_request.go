package certificate_request

import "time"

type CertificateRequest struct {
	Name      string
	Namespace string
	Ready     bool
	Message   string
	Created   time.Time
}
