package SellixGoSDK

import (
	"encoding/json"
	"fmt"
)

type WhitelistObject struct {
	Id        int    `json:"id"`
	UniqueId  string `json:"uniqid"`
	ShopId    int    `json:"shop_id"`
	Type      string `json:"type"`
	Data      string `json:"data"`
	Note      string `json:"note"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	UpdatedBy int    `json:"updated_by"`
}

type GetWhitelistResponseType struct {
	SellixResponseType
	Data *WhitelistObject
}

func (c *SellixClient) GetWhitelist(whitelistId string) (*GetWhitelistResponseType, error) {
	body, err := c.createRequest("GET", "/whitelists/"+whitelistId, nil)
	if err != nil {
		return nil, err
	}
	var sellixResponse GetWhitelistResponseType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

type ListWhitelistResponseType struct {
	SellixResponseType
	Data []WhitelistObject
}

func (c *SellixClient) ListWhitelist() (*ListWhitelistResponseType, error) {
	body, err := c.createRequest("GET", "/whitelists", nil)
	if err != nil {
		return nil, err
	}
	var sellixResponse ListWhitelistResponseType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

type WhitelistEntry struct {
	Type string `json:"type"`
	Data string `json:"data"`
	Note string `json:"note"`
}

func validateWhitelistType(t string) bool {
	validTypes := map[string]struct{}{
		"EMAIL":   {},
		"IP":      {},
		"COUNTRY": {},
		"ISP":     {},
		"ASN":     {},
		"HOST":    {},
	}
	_, ok := validTypes[t]
	return ok
}

type CreateWhitelistResponseType struct {
	SellixResponseType
	Data *struct {
		UniqueId string `json:"uniqid"`
	}
}

func (c *SellixClient) CreateWhitelist(p *WhitelistEntry) (*CreateWhitelistResponseType, error) {
	ok := validateWhitelistType(p.Type)
	if !ok {
		return nil, fmt.Errorf("invalid whitelist type: %s", p.Type)
	}

	body, err := c.createRequest("POST", "/whitelists", p)
	if err != nil {
		return nil, err
	}

	var sellixResponse CreateWhitelistResponseType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

func (c *SellixClient) UpdateWhitelist(p *WhitelistEntry) (*SellixResponseType, error) {
	ok := validateWhitelistType(p.Type)
	if !ok {
		return nil, fmt.Errorf("invalid whitelist type: %s", p.Type)
	}

	body, err := c.createRequest("PUT", "/whitelists", p)
	if err != nil {
		return nil, err
	}

	var sellixResponse SellixResponseType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

func (c *SellixClient) DeleteWhitelist(whitelistId string) (*SellixResponseType, error) {
	body, err := c.createRequest("DELETE", "/whitelists/"+whitelistId, nil)
	if err != nil {
		return nil, err
	}

	var sellixResponse SellixResponseType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}
