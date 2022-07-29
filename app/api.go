package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func (app *application) call(endpoint string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, app.config.apiurl+endpoint, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
