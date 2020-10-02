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

func ExecuteRequest(url string) (Response, error) {
	var response Response
	req, err := http.NewRequest("GET", url, nil)
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
