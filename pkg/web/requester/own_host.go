package requester

import "net/url"

func (r *Requester) ownHost() string {
	u, e := url.Parse(r.base.String())

	if e != nil {
		return ""
	}

	return u.Host
}
