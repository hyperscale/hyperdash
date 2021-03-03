package bus

import "github.com/hyperscale/hyperdash/pkg/hyperdash/protocol"

// Writer interface
type Writer interface {
	Publish(msg *protocol.Message) error
}

// Reader interface
type Reader interface {
	Retrieve() <-chan *protocol.Message
}

// Bus struct
type Bus struct {
	messages chan *protocol.Message
}

// New Bus struct
func New() *Bus {
	return &Bus{
		messages: make(chan *protocol.Message, 500),
	}
}

// Publish msg to bus
func (b *Bus) Publish(msg *protocol.Message) error {
	b.messages <- msg

	return nil
}

// Retrieve msg from bus
func (b *Bus) Retrieve() <-chan *protocol.Message {
	return b.messages
}
