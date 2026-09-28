package band

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (c *Client) envelope(
	action string,
	resource string,
	selectors string,
	body string,
) string {
	return fmt.Sprintf(
		`<?xml version="1.0" encoding="utf-8"?><s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:w="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><s:Header><a:To>%s</a:To><w:ResourceURI s:mustUnderstand="true">%s</w:ResourceURI><a:ReplyTo><a:Address s:mustUnderstand="true">%s</a:Address></a:ReplyTo><a:Action s:mustUnderstand="true">%s</a:Action><w:MaxEnvelopeSize s:mustUnderstand="true">51200</w:MaxEnvelopeSize><a:MessageID>%s</a:MessageID><w:OperationTimeout>PT60S</w:OperationTimeout>%s</s:Header><s:Body>%s</s:Body></s:Envelope>`,
		join.Empty(c.base, constant.Path),
		resource,
		constant.AnonymousRole,
		action,
		messageIdentifier(),
		selectors,
		body,
	)
}
