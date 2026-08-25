package scan

type errorHandlingOperation struct {
	OperationIdentifier string                           `yaml:"operationId"`
	Responses           map[string]errorHandlingResponse `yaml:"responses"`
}
