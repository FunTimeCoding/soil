package plan

type Plan struct {
	ResourceChanges []*ResourceChange `json:"resource_changes"`
	Configuration   *Configuration    `json:"configuration"`
	PlannedValues   *Values           `json:"planned_values"`
	PriorState      *State            `json:"prior_state"`
}
