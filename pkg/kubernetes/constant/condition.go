package constant

const (
	StatusField            = "status"
	ConditionsField        = "conditions"
	SpecField              = "spec"
	IssuerReferenceField   = "issuerRef"
	IssuerReferenceName    = "name"

	ConditionFieldType       = "type"
	ConditionFieldStatus     = "status"
	ConditionFieldMessage    = "message"
	ConditionFieldReason     = "reason"
	ConditionFieldTransition = "lastTransitionTime"

	ConditionReady      = "Ready"
	ConditionIssuing    = "Issuing"
	ConditionStatusTrue = "True"

	ConditionReasonTriggered  = "ManuallyTriggered"
	ConditionMessageTriggered = "Certificate re-issuance triggered"
)
