package SellixGoSDK

import (
	"encoding/json"
	"fmt"
)

type BlacklistType struct {
	Status int `json:"status"`
	Data   struct {
		Blacklist struct {
			Id        int    `json:"id"`
			UniqueId  string `json:"uniqid"`
			Scope     string `json:"scope"`
			ShopId    int    `json:"shop_id"`
			Type      string `json:"type"`
			Data      string `json:"data"`
			Note      string `json:"note"`
			CreatedAt int64  `json:"created_at"`
			UpdatedAt int64  `json:"updated_at"`
			UpdatedBy int    `json:"updated_by"`
		} `json:"blacklist"`
	} `json:"data"`
	Message interface{} `json:"message"`
	Log     interface{} `json:"log"`
	Error   interface{} `json:"error"`
	Env     string      `json:"env"`
}

func (c *SellixClient) GetBlacklist(blacklistId string) (*BlacklistType, error) {
	body, err := c.createRequest("GET", "/blacklists/"+blacklistId, nil)
	if err != nil {
		return nil, err
	}

	var sellixResponse BlacklistType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

func (c *SellixClient) ListBlacklists() (*BlacklistType, error) {
	body, err := c.createRequest("GET", "/blacklists", nil)
	if err != nil {
		return nil, err
	}

	var sellixResponse BlacklistType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

type BlacklistEntry struct {
	Type string `json:"type"`
	Data string `json:"data"`
	Note string `json:"note"`
}

func validateBlacklistType(t string) bool {
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

func (c *SellixClient) CreateBlacklist(p *BlacklistEntry) (*BlacklistType, error) {
	ok := validateBlacklistType(p.Type)
	if !ok {
		return nil, fmt.Errorf("invalid Blacklist Type: %s", p.Type)
	}

	body, err := c.createRequest("POST", "/blacklists", p)
	if err != nil {
		return nil, err
	}

	var sellixResponse BlacklistType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

func (c *SellixClient) UpdateBlacklist(p *BlacklistEntry) (*BlacklistType, error) {
	ok := validateBlacklistType(p.Type)
	if !ok {
		return nil, fmt.Errorf("invalid Blacklist Type: %s", p.Type)
	}

	body, err := c.createRequest("PUT", "/blacklists", p)
	if err != nil {
		return nil, err
	}

	var sellixResponse BlacklistType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

func (c *SellixClient) DeleteBlacklist(blacklistId string) (*BlacklistType, error) {
	body, err := c.createRequest("DELETE", "/blacklists/"+blacklistId, nil)
	if err != nil {
		return nil, err
	}

	var sellixResponse BlacklistType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}
