package policy

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/distinguished_name"
)

func (p *Policy) Validate(n *distinguished_name.Name) {
	errors.PanicOnEmpty(n.CommonName, constant.CommonNameLabel)

	if p.RequireCountry {
		errors.PanicOnEmpty(n.Country, constant.CountryLabel)
	}

	if p.RequireProvince {
		errors.PanicOnEmpty(n.Province, constant.ProvinceLabel)
	}

	if p.RequireOrganization {
		errors.PanicOnEmpty(n.Organization, constant.OrganizationLabel)
	}
}
