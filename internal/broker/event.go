package broker

import "github.com/ozoncp/ocp-certificate-api/internal/model"

type ActionType int

const (
	Create ActionType = iota
	Update
	Remove
	MultiCreate
)

// Message - struct to build messages
type Message struct {
	Type ActionType
	Body model.CertificateEvent
}

// AsyncMultipleCreate - struct to build messages
type AsyncMultipleCreate struct {
	Type ActionType
	Body []model.Certificate
}

// CreateMessage - build messages and send to kafka
func CreateMessage(actionType ActionType, certificateID model.CertificateEvent) Message {
	return Message{
		Type: actionType,
		Body: certificateID,
	}
}

// CreateMessages - build messages and send to kafka
func CreateMessages(actionType ActionType, certificates []model.Certificate) AsyncMultipleCreate {
	return AsyncMultipleCreate{
		Type: actionType,
		Body: certificates,
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
	case MultiCreate:
		return "MultiCreate"
	default:
		return "Unknown MessageType"
	}
}
