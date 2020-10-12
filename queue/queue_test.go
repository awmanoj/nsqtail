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

func TestEnqueue(t *testing.T) {
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