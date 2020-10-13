package queue

import "testing"

var listOfMaxMessagesValues = []int{1000000, 10, 0}
var usualMaxMessageValue = 10

func TestNewInMemoryMessageQueue(t *testing.T) {
	for _, maxMessages := range listOfMaxMessagesValues {
		messageQueue := NewInMemoryMessageQueue(maxMessages)
		if messageQueue == nil {
			t.Errorf("object creation failed, maxMessages: [%d]\n", maxMessages)
		}

		if len(messageQueue.messages) != maxMessages {
			t.Errorf("length of messages different. expected %d actual %d", maxMessages, len(messageQueue.messages))
		}
	}
}

func TestInMemoryMessageQueue_Enqueue(t *testing.T) {
	var testMessages = []string{
		"Hello",
		"Hello again",
		"This",
		"is a short string",
		"and this one is quite loooooooooooooooooooooooooooooooooooooooong",
		"{'this': 'is', 'a': 'json', 'with_strings_and_numbers': 10 }",
		"seventh message",
		"eightth message",
		"ninth message",
		"tenth message",
	}

	var overflowMessages = []string{
		"eleventh message",
	}

	messageQueue := NewInMemoryMessageQueue(usualMaxMessageValue)

	// fill the queue
	for _, testMessage := range testMessages {
		err := messageQueue.Enqueue(testMessage)
		if err != nil {
			t.Errorf("error enqueuing message to the queue %s", err.Error())
		}
	}

	// these will overwrite the oldest messages
	for _, overflowMessage := range overflowMessages {
		err := messageQueue.Enqueue(overflowMessage)
		if err != nil {
			t.Errorf("error enqueuing message to the queue %s", err.Error())
		}
	}

	// hence total count remains the same
	if messageQueue.itemCount != len(testMessages) {
		t.Errorf("error diff count of messages in the queue and the requests. actual [%d] expected [%d]\n",
			messageQueue.itemCount, len(testMessages))
	}

}

func TestInMemoryMessageQueue_Dequeue(t *testing.T) {
	var testMessages = []string{
		"Hello",
		"Hello again",
		"This",
		"is a short string",
		"and this one is quite loooooooooooooooooooooooooooooooooooooooong",
		"{'this': 'is', 'a': 'json', 'with_strings_and_numbers': 10 }",
		"seventh message",
		"eightth message",
		"ninth message",
		"tenth message",
	}

	messageQueue := NewInMemoryMessageQueue(usualMaxMessageValue)

	// try dequeue on empty MQ
	_, err := messageQueue.Dequeue()
	if err == nil {
		t.Errorf("error expected error dequeuing message to the queue")
	}

	// fill the queue
	for _, testMessage := range testMessages {
		err := messageQueue.Enqueue(testMessage)
		if err != nil {
			t.Errorf("error enqueuing message to the queue %s", err.Error())
		}
	}

	// hence total count remains the same
	if messageQueue.itemCount != len(testMessages) {
		t.Errorf("error diff count of messages in the queue and the requests. actual [%d] expected [%d]\n",
			messageQueue.itemCount, len(testMessages))
	}

	for ;messageQueue.itemCount!=0; {
		_, err := messageQueue.Dequeue()
		if err != nil {
			t.Errorf("error dequeuing message to the queue %s", err.Error())
		}
	}
	// try dequeue on empty MQ
	_, err = messageQueue.Dequeue()
	if err == nil {
		t.Errorf("error expected error dequeuing message to the queue")
	}
}

func TestInMemoryMessageQueue_Snapshot(t *testing.T) {
	var input = []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"0",
	}

	messageQueue := NewInMemoryMessageQueue(usualMaxMessageValue)

	// fill the queue
	for _, testMessage := range input {
		err := messageQueue.Enqueue(testMessage)
		if err != nil {
			t.Errorf("error enqueuing message to the queue %s", err.Error())
		}
	}

	output := messageQueue.Snapshot()

	if len(output) != len(input) {
		t.Errorf("discrepancy between the length of input and output")
	}

	for i:=0; i<len(input); i++ {
		if input[i] != output[i] {
			t.Errorf("input[%s] != output[%s]", input[i], output[i])
		}
	}
}

func TestInMemoryMessageQueue_IsEmpty(t *testing.T) {
	messageQueue := NewInMemoryMessageQueue(usualMaxMessageValue)

	if !messageQueue.IsEmpty() {
		t.Errorf("message queue is empty but isEmpty returning false")
	}
}

func TestInMemoryMessageQueue_IsFull(t *testing.T) {
	var input = []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"0",
	}

	messageQueue := NewInMemoryMessageQueue(usualMaxMessageValue)

	// fill the queue
	for _, testMessage := range input {
		err := messageQueue.Enqueue(testMessage)
		if err != nil {
			t.Errorf("error enqueuing message to the queue %s", err.Error())
		}
	}

	if !messageQueue.IsFull() {
		t.Errorf("message queue is full but isFull returning false")
	}
}

func TestInMemoryMessageQueue_Peek(t *testing.T) {
	var input = []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"0",
	}

	messageQueue := NewInMemoryMessageQueue(usualMaxMessageValue)

	// fill the queue
	for _, testMessage := range input {
		err := messageQueue.Enqueue(testMessage)
		if err != nil {
			t.Errorf("error enqueuing message to the queue %s", err.Error())
		}
	}

	value, err := messageQueue.Peek()
	if err != nil {
		t.Errorf("error in peek")
	}

	if value != "1" {
		t.Errorf("wrong peek value")
	}
}

