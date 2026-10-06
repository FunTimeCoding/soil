package open_api

type Operation struct {
	OperationIdentifier string              `yaml:"operationId"`
	Responses           map[string]Response `yaml:"responses"`
}
