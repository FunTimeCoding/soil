package constant

const (
	DeliveryBudget     = 2000
	DeliveryCutFloor   = 300
	DeliveryEntryLimit = 500
)

const (
	DeliverySessionActivity = "Session activity:"
	DeliveryMemoryActivity  = "Memory activity:"
	DeliveryMessages        = "Messages:"
	DeliveryNotifications   = "Notifications:"
	DeliveryIndent          = "  %s"
	DeliveryPulse           = "[pulse] %s"
	DeliveryIdle            = "Idle: %s"
)

const (
	DeliveryExplanation = "This delivery holds about 2,000 characters; longer messages are cut or waiting, and read_message returns them whole."
	DeliveryCut         = "…cut here - %s more characters in message %d, read it with read_message"
	DeliveryPointer     = "Message %d from %s at %s (%s characters) is waiting - read it with read_message"
	DeliveryMemoryTrim  = "Memory changes not shown: %d - list_memories lists recent updates"
)

const (
	ReadMessage       = "read_message"
	Identifiers       = "identifiers"
	MessageHeader     = "[Message %d from %s to %s at %s]"
	MessageEveryone   = "everyone"
	MessageNotFound   = "No message with identifier: %s"
)

const (
	DeliveryOverBudget = "hook delivery over budget"
	DeliveryContextKey = "delivery"
	DeliveryCharacters = "characters"
	DeliveryRefused    = "%s is %d characters; summarize it to %d or fewer"
)
