package SellixGoSDK

import "net/http"

type SellixClient struct {
	Http      *http.Client
	Request   *http.Request
	ApiKey    string
	storeName *string
	baseUrl   string
}

type SellixResponseType struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Log     interface{} `json:"log"`
	Error   interface{} `json:"error"`
	Env     string      `json:"env"`
}

type PostRequestSellixResponse struct {
	SellixResponseType
	Data *struct {
		UniqueId string `json:"uniqid"`
	}
}
