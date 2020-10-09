package nsq

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

var LookupdAddrEnv = "NSQ_LOOKUPD_ADDRESS"

func baseURL() string {
	return "http://" + os.Getenv(LookupdAddrEnv)
}

func getEndpoint(endpoint string) string {
	return baseURL() + endpoint
}


// POST /channel/create?topic=XXXXX&channel=YYYYYY
func CreateChannel(topic string, channel string) error {
	args := map[string]string{
		topic: topic,
		channel: channel,
	}

	response, err := ExecutePostNetworkRequest(getEndpoint("/channel/create"), args)
	if err != nil {
		return err
	}

	if response.statusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("error in the network request to lookupd %d", response.statusCode))
	}

	return nil
}

// POST /channel/delete?topic=XXXXX&channel=YYYYYY
func DeleteChannel(topic string, channel string) error {
	args := map[string]string{
		topic: topic,
		channel: channel,
	}

	response, err := ExecutePostNetworkRequest(getEndpoint("/channel/delete"), args)
	if err != nil {
		return err
	}

	if response.statusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("error in the network request to lookupd %d", response.statusCode))
	}

	return nil
}

// GET /channels
func GetChannels(topic string) (Channels, error) {
	var channels Channels

	response, err := ExecuteGetNetworkRequest(getEndpoint("/channels"), nil)
	if err != nil {
		return channels, err
	}

	if response.statusCode != http.StatusOK {
		return channels, errors.New(fmt.Sprintf("error in the network request to lookupd %d", response.statusCode))
	}

	err = json.Unmarshal(response.body, channels)
	if err != nil {
		return channels, err
	}
	return channels, nil
}

// GET /topics
func GetTopics() (Topics, error) {
	var topics Topics

	response, err := ExecuteGetNetworkRequest(getEndpoint("/topics"), nil)
	if err != nil {
		return topics, err
	}

	if response.statusCode != http.StatusOK {
		return topics, errors.New(fmt.Sprintf("error in the network request to lookupd %d", response.statusCode))
	}

	err = json.Unmarshal(response.body, topics)
	if err != nil {
		return topics, err
	}

	return topics, nil
}
