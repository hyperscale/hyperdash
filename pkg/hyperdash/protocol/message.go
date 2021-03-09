package protocol

// MessageType type
type MessageType string

// MessageType enums
const (
	MessageTypeStat   MessageType = "stat"
	MessageTypeList   MessageType = "list"
	MessageTypeMap    MessageType = "map"
	MessageTypeStatus MessageType = "status"
	MessageTypeUpdate MessageType = "update"
	MessageTypeError  MessageType = "error"
)

// StatusType type
type StatusType string

// StatusType enums
const (
	StatusTypeGreen  StatusType = "green"
	StatusTypeYellow StatusType = "yellow"
	StatusTypeRed    StatusType = "red"
)

// ErrorLevelType type
type ErrorLevelType string

// ErrorType enums
const (
	ErrorLevelTypeError   ErrorLevelType = "error"
	ErrorLevelTypeWarning ErrorLevelType = "warning"
)

// Message struct
type Message struct {
	Target  string      `json:"target"`
	Type    MessageType `json:"type"`
	Payload interface{} `json:"payload"`
}

// MessageStatus struct
type MessageStatus struct {
	Status StatusType `json:"status"`
}

// MessageStat struct
type MessageStat struct {
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

// MessageError struct
type MessageError struct {
	Level ErrorLevelType `json:"level"`
}
