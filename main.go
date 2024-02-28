package SellixGoSDK

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func NewSellixClient(apiKey string) *SellixClient {
	return &SellixClient{
		Http:    &http.Client{},
		ApiKey:  apiKey,
		baseUrl: "https://dev.sellix.io/v1",
	}
}

func (c *SellixClient) SetStoreName(storeName string) {
	c.storeName = &storeName
}

func (c *SellixClient) createRequest(method string, url string, p interface{}) ([]byte, error) {
	payload, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.baseUrl, url), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	req.Header.Set("X-Sellix-Merchant", *c.storeName)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
