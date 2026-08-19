package distinguished_name

import "crypto/x509/pkix"

func From(s pkix.Name) *Name {
	n := New()
	n.Country = first(s.Country)
	n.Province = first(s.Province)
	n.Locality = first(s.Locality)
	n.Organization = first(s.Organization)
	n.Unit = first(s.OrganizationalUnit)
	n.CommonName = s.CommonName

	return n
}
