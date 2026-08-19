package service

import "crypto/x509"

func host(r *x509.CertificateRequest) []string {
	result := r.DNSNames

	for _, a := range r.IPAddresses {
		result = append(result, a.String())
	}

	return result
}
