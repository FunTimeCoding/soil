package technitium

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/technitium/record"
	"net/url"
)

func (c *Client) AddRecord(
	domain string,
	recordType string,
	value string,
) (*record.Record, error) {
	var result addRecordResponse
	path := fmt.Sprintf(
		"/zones/records/add?domain=%s&type=%s",
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

	return result.AddedRecord, c.get(path, &result)
}
