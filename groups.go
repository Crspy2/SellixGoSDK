package SellixGoSDK

import (
	"encoding/json"
)

type GroupObject struct {
	Id            int             `json:"id"`
	Uniqid        string          `json:"uniqid"`
	ShopId        int             `json:"shop_id"`
	Title         string          `json:"title"`
	Unlisted      bool            `json:"unlisted"`
	SortPriority  int             `json:"sort_priority"`
	ProductsBound []ProductObject `json:"products_bound"`
	ProductsCount int             `json:"products_count"`
	CreatedAt     int64           `json:"created_at"`
	UpdatedAt     int64           `json:"updated_at"`
	UpdatedBy     int             `json:"updated_by"`
}

type GetGroupResponseType struct {
	SellixResponseType
	Data *GroupObject
}

func (c *SellixClient) GetGroup(groupId string) (*GetGroupResponseType, error) {
	body, err := c.createRequest("GET", "/groups/"+groupId, nil)
	if err != nil {
		return nil, err
	}
	var sellixResponse GetGroupResponseType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

type ListGroupsResponseType struct {
	SellixResponseType
	Data []GroupObject
}

func (c *SellixClient) ListGroups() (*ListGroupsResponseType, error) {
	body, err := c.createRequest("GET", "/groups", nil)
	if err != nil {
		return nil, err
	}

	var sellixResponse ListGroupsResponseType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

type GroupInfo struct {
	Title         string   `json:"title"`
	Unlisted      bool     `json:"unlisted"`
	ProductsBound []string `json:"products_bound"`
	SortPriority  int      `json:"sort_priority"`
}

func (c *SellixClient) CreateGroup(p *GroupInfo) (*SellixResponseType, error) {
	body, err := c.createRequest("POST", "/groups", p)
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

func (c *SellixClient) EditGroup(groupId string, p *GroupInfo) (*SellixResponseType, error) {
	body, err := c.createRequest("PUT", "/groups/"+groupId, p)
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

func (c *SellixClient) DeleteGroup(groupId string) (*SellixResponseType, error) {
	body, err := c.createRequest("DELETE", "/groups/"+groupId, nil)
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
