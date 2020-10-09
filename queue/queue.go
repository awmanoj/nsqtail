package queue

import "errors"

type InMemoryMessageQueue struct {
	maxSize int
	messages []string

	front int
	rear int
	itemCount int

	enableEvictionOnFull bool
}

func NewInMemoryMessageQueue(maxSize int) *InMemoryMessageQueue {
	return &InMemoryMessageQueue{
		maxSize: maxSize,
		messages: make([]string, maxSize),
		front: 0,
		rear: -1,
		itemCount: 0,
		enableEvictionOnFull: true,
	}
}

func (h *InMemoryMessageQueue) Enqueue(message string) error {
	if h.IsFull() {
		if h.enableEvictionOnFull {
			h.Dequeue()
		} else {
			return errors.New("queue is full") // overflow
		}
	}

	if h.rear == h.maxSize - 1 {
		h.rear = -1
	}

	h.rear++
	h.messages[h.rear] = message
	h.itemCount++

	return nil
}

func (h *InMemoryMessageQueue) Dequeue() (string, error) {
	if h.IsEmpty() {
		return "", errors.New("queue is empty")
	}

	message := h.messages[h.front]
	h.front++

	if h.front == h.maxSize {
		h.front = 0
	}

	h.itemCount--

	return message, nil
}

func (h *InMemoryMessageQueue) Peek() (string, error) {
	if h.IsEmpty() {
		return "", errors.New("queue is empty")
	}

	return h.messages[h.front], nil
}

func (h InMemoryMessageQueue) IsFull() bool {
	return h.itemCount == h.maxSize
}

func (h InMemoryMessageQueue) IsEmpty() bool {
	return h.itemCount == 0
}

// Antipattern - going through the queue and serialize it
// not supposed to be like this but..
func (h *InMemoryMessageQueue) Snapshot() []string {
	var messages []string
	for i := 0; i < h.itemCount; i++ {
		messages = append(messages, h.messages[i])
	}
	return messages
}