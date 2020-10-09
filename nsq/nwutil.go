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

func ExecutePostNetworkRequest(URL string, args map[string]string) (Response, error) {
	return ExecuteNetworkRequest("POST", URL, args)
}

func ExecuteGetNetworkRequest(URL string, args map[string]string) (Response, error) {
	return ExecuteNetworkRequest("GET", URL, args)
}

func ExecuteNetworkRequest(method string, URL string, args map[string]string) (Response, error) {
	if len(args) > 0 {
		URL += "?"
	}

	for key,value := range args {
		URL += key + "=" + value
	}

	var response Response
	req, err := http.NewRequest(method, URL, nil)
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
