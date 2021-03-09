package bus

import (
	"fmt"
	"sync"

	"github.com/hyperscale/hyperdash/pkg/hyperdash/protocol"
	"github.com/rs/zerolog/log"
)

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
	messages     chan *protocol.Message
	receivers    map[string]chan *protocol.Message
	receiversMtx sync.RWMutex
	done         chan bool
}

// New Bus struct
func New() *Bus {
	return &Bus{
		messages:  make(chan *protocol.Message, 500),
		receivers: make(map[string]chan *protocol.Message),
		done:      make(chan bool),
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

func (b *Bus) Start() {
	for {
		select {
		case t := <-b.messages:
			// breadcast
			for _, reciever := range b.getReceivers() {
				reciever <- t
			}
		case <-b.done:
			log.Debug().Msg("Done")

			return
		}
	}
}

func (b *Bus) Stop() {
	b.done <- true
}

func (b *Bus) getReceivers() []chan *protocol.Message {
	receivers := []chan *protocol.Message{}

	b.receiversMtx.RLock()
	defer b.receiversMtx.RUnlock()

	for _, receiver := range b.receivers {
		receivers = append(receivers, receiver)
	}

	return receivers
}

// Register receiver
func (b *Bus) Register(id string, ch chan *protocol.Message) error {
	b.receiversMtx.Lock()
	defer b.receiversMtx.Unlock()

	if _, ok := b.receivers[id]; ok {
		return fmt.Errorf("reciever id %q is already registered", id)
	}

	b.receivers[id] = ch

	return nil
}

// Unregister receiver
func (b *Bus) Unregister(id string) {
	b.receiversMtx.Lock()
	defer b.receiversMtx.Unlock()

	delete(b.receivers, id)
}
