package connection

import "github.com/funtimecoding/soil/pkg/assistant/message"

type Subscriber func(*message.Message)
