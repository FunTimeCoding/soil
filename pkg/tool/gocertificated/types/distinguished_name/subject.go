package distinguished_name

import "crypto/x509/pkix"

func (n *Name) Subject() pkix.Name {
	s := pkix.Name{}
	s.CommonName = n.CommonName

	if n.Country != "" {
		s.Country = []string{n.Country}
	}

	if n.Province != "" {
		s.Province = []string{n.Province}
	}

	if n.Locality != "" {
		s.Locality = []string{n.Locality}
	}

	if n.Organization != "" {
		s.Organization = []string{n.Organization}
	}

	if n.Unit != "" {
		s.OrganizationalUnit = []string{n.Unit}
	}

	return s
}
