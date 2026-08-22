package technitium

import (
	"fmt"
	"net/url"
)

func (c *Client) DeleteRecord(
	domain string,
	recordType string,
	value string,
) error {
	path := fmt.Sprintf(
		"/zones/records/delete?domain=%s&type=%s",
		url.QueryEscape(domain),
		url.QueryEscape(recordType),
	)

	switch recordType {
	case "A", "AAAA":
		path = fmt.Sprintf("%s&ipAddress=%s", path, url.QueryEscape(value))
	case "CNAME":
		path = fmt.Sprintf("%s&cname=%s", path, url.QueryEscape(value))
	case "PTR":
		path = fmt.Sprintf("%s&ptrName=%s", path, url.QueryEscape(value))
	case "TXT":
		path = fmt.Sprintf("%s&text=%s", path, url.QueryEscape(value))
	}

	_, e := c.basic.Get(path)

	return e
}
