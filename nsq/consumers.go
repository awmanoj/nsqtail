package nsq

import (
	"github.com/nsqio/go-nsq"
	"github.com/koding/cache"
	"log"
)

const MaxNumOfMessages = 10
const channelName = "nsqtail-taxicab-1729" // nsqtail-taxicab-1729 - unique enough?

var consumers = map[string]*nsq.Consumer{}
var caches = map[string]cache.Cache{}

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
		caches[topic] = cache.NewLRU(MaxNumOfMessages)
	}
}

func FetchLastNRequests(topic string, n int) ([]string, error) {
	var lastNRequests []string

	// Right now, only supports maxNumOfMessages requests
	if n > MaxNumOfMessages {
		n = MaxNumOfMessages
	}

	return lastNRequests, nil
}

