package store

import "net/url"

func marshalForm(form url.Values) string {
	return form.Encode()
}
