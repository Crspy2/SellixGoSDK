package sellix

import (
	"encoding/json"
	"fmt"
)

type BlacklistObject struct {
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
}

type GetBlacklistResponseType struct {
	SellixResponseType
	Data *BlacklistObject
}

func (c *SellixClient) GetBlacklist(blacklistId string) (*GetBlacklistResponseType, error) {
	body, err := c.createRequest("GET", "/blacklists/"+blacklistId, nil)
	if err != nil {
		return nil, err
	}
	var sellixResponse GetBlacklistResponseType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

type ListBlacklistResponseType struct {
	SellixResponseType
	Data []BlacklistObject
}

func (c *SellixClient) ListBlacklists() (*ListBlacklistResponseType, error) {
	body, err := c.createRequest("GET", "/blacklists", nil)
	if err != nil {
		return nil, err
	}

	var sellixResponse ListBlacklistResponseType
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

func (c *SellixClient) CreateBlacklist(p *BlacklistEntry) (*PostRequestSellixResponse, error) {
	ok := validateBlacklistType(p.Type)
	if !ok {
		return nil, fmt.Errorf("invalid Blacklist Type: %s", p.Type)
	}

	body, err := c.createRequest("POST", "/blacklists", p)
	if err != nil {
		return nil, err
	}

	var sellixResponse PostRequestSellixResponse
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

func (c *SellixClient) UpdateBlacklist(p *BlacklistEntry) (*SellixResponseType, error) {
	ok := validateBlacklistType(p.Type)
	if !ok {
		return nil, fmt.Errorf("invalid Blacklist Type: %s", p.Type)
	}

	body, err := c.createRequest("PUT", "/blacklists", p)
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

func (c *SellixClient) DeleteBlacklist(blacklistId string) (*SellixResponseType, error) {
	body, err := c.createRequest("DELETE", "/blacklists/"+blacklistId, nil)
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
