package example

var (
	url             = "test"
	locator         = "test" // ok: uses the suggestion
	dirName         = "test"
	directory       = "test" // ok: contains the suggestion
	DirSomething    = "test"
	OutputDirectory = "" // ok: segment is "directory"
	ctx             = ""
	msgText         = ""
	paramCount      = 0
	parameterCount  = 0 // ok: contains the suggestion
	mcpServer       = ""
	clean           = "" // ok: no blacklisted segment
	req             = ""
	request         = "" // ok: uses the suggestion
	docPath         = ""
	document        = "" // ok: uses the suggestion
	cfg             = ""
	llm             = ""
	tmp             = ""
	loginHandler    = ""
	userInfo        = ""
	userData        = ""
	posX            = 0
	position        = 0 // ok: uses the suggestion
	bufData         = []byte{}
	buffer          = []byte{} // ok: uses the suggestion
	ptrValue        = 0
	addrStr         = ""
	refCount        = 0
	navBar          = ""
	prevState       = ""
	past            = "" // ok: uses the suggestion
	yamlData        = []byte{}
	xmlContent      = []byte{}
	htmlPage        = ""
	jsonPayload     = []byte{}
	declCount       = 0
)
