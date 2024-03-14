package sellix

import "encoding/json"

type LicenseObject struct {
	CreatedAt  string `json:"created_at"`
	ExpiresAt  string `json:"expires_at"`
	HardwareId string `json:"hardware_id"`
	Id         int    `json:"id"`
	InvoiceId  string `json:"invoice_id"`
	LicenseKey string `json:"license_key"`
	ProductId  int    `json:"product_id"`
	ShopId     int    `json:"shop_id"`
	UniqueId   string `json:"uniqid"`
	UpdatedAt  string `json:"updated_at"`
	Uses       int    `json:"uses"`
}

type LicenseInfo struct {
	HardwareId string `json:"hardware_id"`
	Key        string `json:"key"`
	ProductId  string `json:"product_id"`
}

func (c *SellixClient) CheckLicense(l LicenseInfo) (*GetOrderResponseType, error) {
	body, err := c.createRequest("POST", "/products/licensing/check", l)
	if err != nil {
		return nil, err
	}
	var sellixResponse GetOrderResponseType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

func (c *SellixClient) SetHardwareID(l LicenseInfo) (*GetOrderResponseType, error) {
	body, err := c.createRequest("PUT", "/products/licensing/hardware_id", l)
	if err != nil {
		return nil, err
	}
	var sellixResponse GetOrderResponseType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}
