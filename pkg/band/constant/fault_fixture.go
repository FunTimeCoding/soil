package constant

const (
	FixtureChallenge = `Digest realm="Digest:A3829B3827DE4D33D4449B366831FD01", nonce="Q3h9JgAAAAAAAAAA", stale="false", qop="auth"`

	FixtureAccessDeniedFault = `<?xml version="1.0" encoding="UTF-8"?><a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" xmlns:b="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><a:Body><a:Fault><a:Code><a:Value>a:Sender</a:Value><a:Subcode><a:Value>b:AccessDenied</a:Value></a:Subcode></a:Code><a:Reason><a:Text xml:lang="en-US">The sender was not authorized to access the resource.</a:Text></a:Reason></a:Fault></a:Body></a:Envelope>`

	FixtureInvalidParameterFault = `<?xml version="1.0" encoding="UTF-8"?><a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" xmlns:b="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><a:Body><a:Fault><a:Code><a:Value>a:Sender</a:Value><a:Subcode><a:Value>b:InvalidParameter</a:Value></a:Subcode></a:Code><a:Reason><a:Text xml:lang="en-US">An operation parameter is not valid.</a:Text></a:Reason></a:Fault></a:Body></a:Envelope>`

	FixtureTimedOutFault = `<?xml version="1.0" encoding="UTF-8"?><a:Envelope xmlns:a="http://www.w3.org/2003/05/soap-envelope" xmlns:b="http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"><a:Body><a:Fault><a:Code><a:Value>a:Receiver</a:Value><a:Subcode><a:Value>b:TimedOut</a:Value></a:Subcode></a:Code><a:Reason><a:Text xml:lang="en-US">The operation has timed out.</a:Text></a:Reason></a:Fault></a:Body></a:Envelope>`
)
