package nsq

import (
	"github.com/nsqio/go-nsq"
	"log"

	"github.com/awmanoj/nsqtail/queue"
)

const MaxNumOfMessages = 10
const channelName = "nsqtail-taxicab-1729" // nsqtail-taxicab-1729 - unique enough?

var consumers = map[string]*nsq.Consumer{}
var queues = map[string]*queue.InMemoryMessageQueue{}

func InitConsumers() {
	topics, err := GetTopics()
	if err != nil {
		log.Printf("err", "problem fetching topics", err)
		return
	}

	for _, topic := range topics.Topics {
		consumer, err := initConsumer(topic, channelName)
		if err != nil {
			log.Println("err", err)
			continue
		}
		consumers[topic] = consumer
		queues[topic] = queue.NewInMemoryMessageQueue(MaxNumOfMessages)
	}
}

func FetchLastNRequests(topic string, n int) ([]string, error) {
	// Right now, only supports maxNumOfMessages requests
	if n > MaxNumOfMessages {
		n = MaxNumOfMessages
	}

	lastNRequests := queues[topic].Snapshot()
	return lastNRequests, nil
}
