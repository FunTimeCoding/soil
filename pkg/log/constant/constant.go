package constant

// Reference: https://opentelemetry.io/docs/specs/semconv/messaging
const (
	MessagingSystem = "messaging.system"
	OperationName   = "messaging.operation.name"

	MessageIdentifier      = "messaging.message.id"
	ConversationIdentifier = "messaging.message.conversation_id"
	BodySize               = "messaging.message.body.size"
	EnvelopeTime           = "messaging.message.envelope_time"

	MessageType = "messaging.message.type"

	HomeAssistantOrigin = "messaging.homeassistant.origin"
)
