package producer

type ActionType int

const (
	Create ActionType = iota
	Update
	Remove
)

// Message - struct to build messages
type Message struct {
	Type ActionType
	Body EventMessage
}

// EventMessage - struct for build broker messages
type EventMessage struct {
	Action    string
	Id        uint64
	Timestamp int64
}

// CreateMessage - build messages and send to kafka
func CreateMessage(actionType ActionType, eventMessage EventMessage) Message {
	return Message{
		Type: actionType,
		Body: eventMessage,
	}
}

// String - convert const to string
func (actionType ActionType) String() string {
	switch actionType {
	case Create:
		return "Created"
	case Update:
		return "Updated"
	case Remove:
		return "Removed"
	default:
		return "Unknown MessageType"
	}
}
