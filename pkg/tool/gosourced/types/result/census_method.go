package result

type CensusMethod struct {
	Name            string          `json:"name"`
	File            string          `json:"file"`
	Callers         []*CensusCaller `json:"callers"`
	Calls           []string        `json:"calls"`
	Functions       []string        `json:"functions"`
	Fields          []string        `json:"fields"`
	Uses            []*CensusUse    `json:"uses"`
	Packages        []string        `json:"packages"`
	Atomically      bool            `json:"atomically"`
	TransactionStep bool            `json:"transaction_step"`
	Tables          *CensusTables   `json:"tables"`
	Reach           *CensusTables   `json:"reach"`
}
