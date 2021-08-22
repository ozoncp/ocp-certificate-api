package producer

import (
	"time"
)

type ActionType int

const (
	Create ActionType = iota
	Update
	Remove
)

// Message - struct to build messages
type Message struct {
	Type ActionType
	Body map[string]interface{}
}

// CreateMessage - build messages and send to kafka
func CreateMessage(actionType ActionType, id uint64) Message {
	return Message{
		Type: actionType,
		Body: map[string]interface{}{
			"Id":        id,
			"Action":    actionType.String(),
			"Timestamp": time.Now().Unix(),
		},
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
