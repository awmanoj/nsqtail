package nsq

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var baseURL = "http://" + instantiated.nsqLookupdAddr

func GetTopics() (Topics, error) {
	var topics Topics
	response, err := ExecuteNetworkRequest(baseURL + "/topics")
	if err != nil {
		return topics, err
	}

	if response.statusCode != http.StatusOK {
		return topics, errors.New(fmt.Sprint("error in the network request to lookupd %d", response.statusCode))
	}

	err = json.Unmarshal(response.body, topics)
	if err != nil {
		return topics, err
	}

	return topics, nil
}
