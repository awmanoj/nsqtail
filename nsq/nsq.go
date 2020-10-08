package nsq

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func baseURL() string {
	return "http://" + instantiated.nsqLookupdAddr
}

func getEndpoint(endpoint string) string {
	return baseURL() + endpoint
}

// GET /channels
func GetChannels(topic string) (Channels, error) {
	var channels Channels

	response, err := ExecuteGetNetworkRequest(getEndpoint("/channels"))
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

	response, err := ExecuteGetNetworkRequest(getEndpoint("/topics"))
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
