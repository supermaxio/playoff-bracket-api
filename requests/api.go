package requests

import (
	"net/http"
	"time"
)

func ApiGetRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	return client.Do(req)
}
