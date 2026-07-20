package constant

import "net/http"

const (
	Success       = "success"
	Unauthorized  = "unauthorized"
	InternalError = "internal error"

	RequestStart = "request_start"

	OriginAll = "*"

	FormMethod     = "method"
	ErrorParameter = "error"

	Extended   = "https://unpkg.com/htmx.org@2.0.4"
	ServerSide = "https://unpkg.com/htmx-ext-sse@2.2.2/sse.js"

	Pico = "https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css"
)

var OkayStatusCodes = []int{http.StatusOK, http.StatusCreated}

const LoginAddressEnvironment = "LOGIN_ADDRESS"
const (
	CertificateName               = "certificate"
	CertificateFileName           = "certificate.pem"
	CertificateKeyName            = "key.pem"
	CertificateFileEnvironment    = "CERTIFICATE_FILE"
	CertificateKeyFileEnvironment = "CERTIFICATE_KEY_FILE"
)
