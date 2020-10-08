package nsq

import (
	"io/ioutil"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: 3 * time.Second,
}

type Response struct {
	statusCode int
	body       []byte
}

func ExecutePostNetworkRequest(url string) (Response, error) {
	return ExecuteNetworkRequest("POST", url)
}

func ExecuteGetNetworkRequest(url string) (Response, error) {
	return ExecuteNetworkRequest("GET", url)
}

func ExecuteNetworkRequest(method string, url string) (Response, error) {
	var response Response
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return response, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()

	response.statusCode = resp.StatusCode
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	response.body = body

	return response, nil
}
