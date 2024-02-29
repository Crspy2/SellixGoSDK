package SellixGoSDK

import (
	"encoding/json"
)

type ProductObject struct {
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
	PriceVariants   []struct {
		Price       int    `json:"price"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Stock       int    `json:"stock"`
	} `json:"price_variants"`
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
	YoutubeLink       string        `json:"youtube_link"`
	WatermarkEnabled  bool          `json:"watermark_enabled"`
	WatermarkText     string        `json:"watermark_text"`
	RedirectLink      string        `json:"redirect_link"`
	Theme             string        `json:"theme"`
	DarkMode          int           `json:"dark_mode"`
	AverageScore      float64       `json:"average_score"`
	SoldCount         int           `json:"sold_count"`
	LexPaymentMethods []interface{} `json:"lex_payment_methods"`
	CreatedAt         int64         `json:"created_at"`
	UpdatedAt         int64         `json:"updated_at"`
	UpdatedBy         int           `json:"updated_by"`
}

type GetProductResponseType struct {
	SellixResponseType
	Data *ProductObject
}

func (c *SellixClient) GetProduct(productId string) (*GetProductResponseType, error) {
	body, err := c.createRequest("GET", "/products/"+productId, nil)
	if err != nil {
		return nil, err
	}
	var sellixResponse GetProductResponseType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

type ListProductsResponseType struct {
	SellixResponseType
	Data []ProductObject
}

func (c *SellixClient) ListProducts() (*ListProductsResponseType, error) {
	body, err := c.createRequest("GET", "/products", nil)
	if err != nil {
		return nil, err
	}

	var sellixResponse ListProductsResponseType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

type ProductInfo struct {
	Title         string   `json:"title"`
	Price         float64  `json:"price"`
	Description   string   `json:"description"`
	Currency      string   `json:"currency"`
	Gateways      []string `json:"gateways"`
	Type          string   `json:"type"`
	Serials       []string `json:"serials"`
	PriceVariants []struct {
		Price       int      `json:"price"`
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Serials     []string `json:"serials"`
		Stock       int      `json:"stock,omitempty"`
	} `json:"price_variants"`
	SerialsRemoveDuplicates   bool        `json:"serials_remove_duplicates"`
	ServiceText               interface{} `json:"service_text"`
	Stock                     interface{} `json:"stock"`
	DynamicWebhook            interface{} `json:"dynamic_webhook"`
	StockDelimiter            string      `json:"stock_delimiter"`
	MinQuantity               int         `json:"min_quantity"`
	MaxQuantity               int         `json:"max_quantity"`
	DeliveryText              string      `json:"delivery_text"`
	CustomFields              interface{} `json:"custom_fields"`
	CryptoConfirmationsNeeded int         `json:"crypto_confirmations_needed"`
	MaxRiskLevel              int         `json:"max_risk_level"`
	Unlisted                  bool        `json:"unlisted"`
	Private                   bool        `json:"private"`
	BlockVpnProxies           bool        `json:"block_vpn_proxies"`
	SortPriority              int         `json:"sort_priority"`
	Webhooks                  interface{} `json:"webhooks"`
	OnHold                    bool        `json:"on_hold"`
	TermsOfService            string      `json:"terms_of_service"`
	Warranty                  int         `json:"warranty"`
	WarrantyText              string      `json:"warranty_text"`
	YoutubeLink               string      `json:"youtube_link"`
	WatermarkEnabled          bool        `json:"watermark_enabled"`
	WatermarkText             string      `json:"watermark_text"`
	RedirectLink              string      `json:"redirect_link"`
	RemoveImages              bool        `json:"remove_images"`
	RemoveFiles               bool        `json:"remove_files"`
	VolumeDiscounts           interface{} `json:"volume_discounts"`
	RecurringInterval         interface{} `json:"recurring_interval"`
	RecurringIntervalCount    interface{} `json:"recurring_interval_count"`
	TrialPeriod               interface{} `json:"trial_period"`
}

func (c *SellixClient) CreateProduct(p *ProductInfo) (*SellixResponseType, error) {
	body, err := c.createRequest("POST", "/products", p)
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

func (c *SellixClient) EditProduct(productId string, p *ProductInfo) (*SellixResponseType, error) {
	body, err := c.createRequest("PUT", "/products/"+productId, p)
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

func (c *SellixClient) DeleteProduct(productId string) (*SellixResponseType, error) {
	body, err := c.createRequest("DELETE", "/products/"+productId, nil)
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
