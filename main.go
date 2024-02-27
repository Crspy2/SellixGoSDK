package SellixGoSDK

import "net/http"

type SellixClient struct {
	Http      *http.Client
	Request   *http.Request
	ApiKey    string
	StoreName *string
}

func NewSellixClient(apiKey string) *SellixClient {
	return &SellixClient{
		Http:   &http.Client{},
		ApiKey: apiKey,
	}
}

func (client *SellixClient) SetStoreName(storeName string) {
	client.StoreName = &storeName
}
