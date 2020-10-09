package nsq

import (
	"errors"
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
)

type messageHandler struct {
	Topic string
}

// HandleMessage implements the Handler interface.
func (h *messageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		// In this case, a message with an empty body is simply ignored/discarded.
		return nil
	}

	// do whatever actual message processing is desired
	err := h.processMessage(m.Body)

	// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
	return err
}

func (h *messageHandler) processMessage(body []byte) error {
	message := string(body) // create a copy

	queues[h.Topic].Enqueue(message)

	return nil
}

func initConsumer(topic string, channel string) (*nsq.Consumer, error) {
	// Instantiate a consumer that will subscribe to the provided channel.
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("problem creating consumer for topic [%s], channel [%s] [%v]",topic, channel, err))
	}

	// Set the Handler for messages received by this Consumer. Can be called multiple times.
	// See also AddConcurrentHandlers.
	consumer.AddHandler(&messageHandler{
		Topic: topic,
	})

	// Use nsqlookupd to discover nsqd instances.
	// See also ConnectToNSQD, ConnectToNSQDs, ConnectToNSQLookupds.
	lookupdAddr := os.Getenv(LookupdAddrEnv)
	err = consumer.ConnectToNSQLookupd(lookupdAddr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("problem connecting to NSQLookupd [%s], for topic [%s], channel [%s] [%v]", lookupdAddr, topic, channel, err))
	}

	return consumer, nil
}

func stopConsumer(consumer *nsq.Consumer) {
	consumer.Stop()
}