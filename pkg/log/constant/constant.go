package constant

// https://opentelemetry.io/docs/specs/semconv/messaging
const (
	// Required
	MessagingSystem = "messaging.system"
	OperationName   = "messaging.operation.name"

	// Recommended
	MessageIdentifier      = "messaging.message.id"
	ConversationIdentifier = "messaging.message.conversation_id"
	BodySize               = "messaging.message.body.size"
	EnvelopeTime           = "messaging.message.envelope_time"

	// MessageType Optional
	MessageType = "messaging.message.type"

	// HomeAssistantOrigin Home Assistant specific
	HomeAssistantOrigin = "messaging.homeassistant.origin"
)
