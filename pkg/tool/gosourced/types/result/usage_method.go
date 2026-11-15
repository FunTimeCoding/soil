package result

type UsageMethod struct {
	Name            string         `json:"name"`
	File            string         `json:"file"`
	Callers         []*UsageCaller `json:"callers"`
	Calls           []string       `json:"calls"`
	Functions       []string       `json:"functions"`
	Fields          []string       `json:"fields"`
	Uses            []*UsedType    `json:"uses"`
	Packages        []string       `json:"packages"`
	Atomically      bool           `json:"atomically"`
	TransactionStep bool           `json:"transaction_step"`
	Tables          *UsageTables   `json:"tables"`
	Reach           *UsageTables   `json:"reach"`
}
