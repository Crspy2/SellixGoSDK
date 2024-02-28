package SellixGoSDK

import (
	"encoding/json"
)

type CategoryObject struct {
	Id            int    `json:"id"`
	UniqueId      string `json:"uniqid"`
	ShopId        int    `json:"shop_id"`
	Title         string `json:"title"`
	Unlisted      bool   `json:"unlisted"`
	SortPriority  int    `json:"sort_priority"`
	ProductsBound []struct {
		Id              int         `json:"id"`
		UniqueId        string      `json:"uniqid"`
		ShopId          int         `json:"shop_id"`
		Type            string      `json:"type"`
		Subtype         interface{} `json:"subtype"`
		Title           string      `json:"title"`
		Currency        string      `json:"currency"`
		Price           int         `json:"price"`
		PriceDisplay    float64     `json:"price_display"`
		Description     string      `json:"description"`
		ImageAttachment interface{} `json:"image_attachment"`
		FileAttachment  interface{} `json:"file_attachment"`
		VolumeDiscounts []struct {
			Type     string `json:"type"`
			Value    int    `json:"value"`
			Quantity int    `json:"quantity"`
		} `json:"volume_discounts"`
		RecurringInterval      string      `json:"recurring_interval"`
		RecurringIntervalCount int         `json:"recurring_interval_count"`
		TrialPeriod            int         `json:"trial_period"`
		PaypalProductId        interface{} `json:"paypal_product_id"`
		PaypalPlanId           interface{} `json:"paypal_plan_id"`
		StripePriceId          string      `json:"stripe_price_id"`
		QuantityMin            int         `json:"quantity_min"`
		QuantityMax            int         `json:"quantity_max"`
		QuantityWarning        int         `json:"quantity_warning"`
		Gateways               []string    `json:"gateways"`
		CustomFields           []struct {
			Type        string      `json:"type"`
			Name        string      `json:"name"`
			Regex       interface{} `json:"regex"`
			Placeholder interface{} `json:"placeholder"`
			Default     interface{} `json:"default"`
			Required    bool        `json:"required"`
		} `json:"custom_fields"`
		CryptoConfirmationsNeeded int           `json:"crypto_confirmations_needed"`
		MaxRiskLevel              int           `json:"max_risk_level"`
		BlockVpnProxies           bool          `json:"block_vpn_proxies"`
		DeliveryText              string        `json:"delivery_text"`
		ServiceText               string        `json:"service_text"`
		StockDelimiter            string        `json:"stock_delimiter"`
		Stock                     int           `json:"stock"`
		DynamicWebhook            interface{}   `json:"dynamic_webhook"`
		SortPriority              int           `json:"sort_priority"`
		Unlisted                  bool          `json:"unlisted"`
		OnHold                    bool          `json:"on_hold"`
		TermsOfService            interface{}   `json:"terms_of_service"`
		Warranty                  int           `json:"warranty"`
		WarrantyText              string        `json:"warranty_text"`
		Private                   bool          `json:"private"`
		Name                      string        `json:"name"`
		ImageName                 interface{}   `json:"image_name"`
		ImageStorage              interface{}   `json:"image_storage"`
		CloudflareImageId         string        `json:"cloudflare_image_id"`
		Serials                   []interface{} `json:"serials"`
		Webhooks                  []interface{} `json:"webhooks"`
		Feedback                  struct {
			Total    int `json:"total"`
			Positive int `json:"positive"`
			Neutral  int `json:"neutral"`
			Negative int `json:"negative"`
		} `json:"feedback"`
		Theme             string        `json:"theme"`
		DarkMode          int           `json:"dark_mode"`
		AverageScore      float64       `json:"average_score"`
		SoldCount         int           `json:"sold_count"`
		LexPaymentMethods []interface{} `json:"lex_payment_methods"`
		CreatedAt         int64         `json:"created_at"`
		UpdatedAt         int64         `json:"updated_at"`
		UpdatedBy         int           `json:"updated_by"`
	} `json:"products_bound"`
	ProductsCount int `json:"products_count"`
	GroupsBound   []struct {
		UniqueId        string      `json:"uniqid"`
		Title           string      `json:"title"`
		ImageAttachment interface{} `json:"image_attachment"`
		CreatedAt       int64       `json:"created_at"`
		UpdatedAt       int64       `json:"updated_at"`
	} `json:"groups_bound"`
	GroupsCount int   `json:"groups_count"`
	CreatedAt   int64 `json:"created_at"`
	UpdatedAt   int64 `json:"updated_at"`
	UpdatedBy   int   `json:"updated_by"`
}

type GetCategoryResponseType struct {
	SellixResponseType
	Data *CategoryObject
}

func (c *SellixClient) GetCategory(categoryId string) (*GetCategoryResponseType, error) {
	body, err := c.createRequest("GET", "/categories/"+categoryId, nil)
	if err != nil {
		return nil, err
	}
	var sellixResponse GetCategoryResponseType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

type ListCategoriesResponseType struct {
	SellixResponseType
	Data []CategoryObject
}

func (c *SellixClient) ListCategories() (*ListCategoriesResponseType, error) {
	body, err := c.createRequest("GET", "/categories", nil)
	if err != nil {
		return nil, err
	}

	var sellixResponse ListCategoriesResponseType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

type CategoryEntry struct {
	Title         string   `json:"title"`
	Unlisted      bool     `json:"unlisted"`
	ProductsBound []string `json:"products_bound"`
	GroupsBound   []string `json:"groups_bound"`
	SortPriority  int      `json:"sort_priority"`
}

func (c *SellixClient) CreateCategory(p *CategoryEntry) (*PostRequestSellixResponse, error) {
	body, err := c.createRequest("POST", "/categories", p)
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

func (c *SellixClient) EditCategory(p *CategoryEntry) (*SellixResponseType, error) {
	body, err := c.createRequest("PUT", "/categories", p)
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

func (c *SellixClient) DeleteCategory(categoryId string) (*SellixResponseType, error) {
	body, err := c.createRequest("DELETE", "/categories/"+categoryId, nil)
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
