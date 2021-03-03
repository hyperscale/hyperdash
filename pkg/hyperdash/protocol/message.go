package protocol

// MessageType type
type MessageType string

// MessageType enums
const (
	MessageTypeSignle MessageType = "signle"
	MessageTypeList   MessageType = "list"
	MessageTypeMap    MessageType = "map"
	MessageTypeStatus MessageType = "status"
)

// StatusType type
type StatusType string

// StatusType enums
const (
	StatusTypeGreen  StatusType = "green"
	StatusTypeYellow StatusType = "yellow"
	StatusTypeRed    StatusType = "red"
)

// Message struct
type Message struct {
	Target  string      `json:"target"`
	Type    MessageType `json:"type"`
	Payload interface{} `json:"payload"`
}
