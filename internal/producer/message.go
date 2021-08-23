package producer

import "time"

type ActionType int

const (
	Create ActionType = iota
	Update
	Remove
)

// Message - struct to build messages
type Message struct {
	Type ActionType
	Body BodyMessage
}

// BodyMessage - struct for build broker messages
type BodyMessage struct {
	Action    string
	Id        uint64
	Timestamp int64
}

// CreateMessage - build messages and send to kafka
func CreateMessage(actionType ActionType, id uint64, time time.Time) Message {
	return Message{
		Type: actionType,
		Body: BodyMessage{
			Action:    actionType.String(),
			Id:        id,
			Timestamp: time.Unix(),
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
