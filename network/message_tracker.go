package network

import (
	"errors"
	"sync"
)

// MessageTracker tracks a configurable fixed amount of messages.
// Messages are stored first-in-first-out.  Duplicate messages should not be stored in the queue.
type MessageTracker interface {
	// Add will add a message to the tracker, deleting the oldest message if necessary
	Add(message *Message) (err error)
	// Delete will delete message from tracker
	Delete(id string) (err error)
	// Get returns a message for a given ID.  Message is retained in tracker
	Message(id string) (message *Message, err error)
	// Messages returns messages in FIFO order
	Messages() (messages []*Message)
}

// ErrMessageNotFound is an error returned by MessageTracker when a message with specified id is not found
var ErrMessageNotFound = errors.New("message not found")

func NewMessageTracker(length int) MessageTracker {
 return &MessageTrackerImpl{
  messages: make(map[string]*Message),
  maxLen:   length,
 }
}

type MessageTrackerImpl struct {
 messages map[string]*Message
 order    []string
 maxLen   int
 mu       sync.Mutex
}

func (mt *MessageTrackerImpl) Add(message *Message) error {
 mt.mu.Lock()
 defer mt.mu.Unlock()

 if _, ok := mt.messages[message.ID]; ok {
  // If the message already exists, just return without error
  return nil
 }

 mt.messages[message.ID] = message
 mt.order = append(mt.order, message.ID)

 if len(mt.order) > mt.maxLen {
  delete(mt.messages, mt.order[0])
  mt.order = mt.order[1:]
 }

 return nil
}

func (mt *MessageTrackerImpl) Delete(id string) error {
 mt.mu.Lock()
 defer mt.mu.Unlock()

 if _, ok := mt.messages[id]; !ok {
  return ErrMessageNotFound
 }

 delete(mt.messages, id)

 for i, v := range mt.order {
  if v == id {
   mt.order = append(mt.order[:i], mt.order[i+1:]...)
   break
  }
 }

 return nil
}

func (mt *MessageTrackerImpl) Message(id string) (*Message, error) {
 mt.mu.Lock()
 defer mt.mu.Unlock()

 if msg, ok := mt.messages[id]; ok {
  return msg, nil
 }

 return nil, ErrMessageNotFound
}

func (mt *MessageTrackerImpl) Messages() []*Message {
 mt.mu.Lock()
 defer mt.mu.Unlock()

 msgs := make([]*Message, 0, len(mt.order))
 for _, v := range mt.order {
  msgs = append(msgs, mt.messages[v])
 }

 return msgs
}

