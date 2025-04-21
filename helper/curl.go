package helper

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"
)

func CreateHttpReq(url, method, token, apiKey, body string) (int, []byte, error) {

	httpRequest, err := http.NewRequest(method, url, bytes.NewReader([]byte(body)))
	if err != nil {
		return 0, nil, err
	}
	httpRequest.Header.Set("Content-Type", "application/json")
	
	if token != "" {
		httpRequest.Header.Set("Authorization", "Bearer "+token)
	}

	if apiKey != "" {
		log.Log("API KEY : ", apiKey)
		httpRequest.Header.Set("X-Api-Key", apiKey)
	}

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	log.Log("Resp Body : ", body)

	res, err := client.Do(httpRequest)
	if err != nil {
		return 0, nil, err
	}

	if res.StatusCode < 200 && res.StatusCode > 299 {
		resBody, _ := io.ReadAll(res.Body)
		return res.StatusCode, resBody, errors.New(strconv.Itoa(res.StatusCode) + " : " + res.Status)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, nil, err
	}
	return res.StatusCode, resBody, nil
}
