package SellixGoSDK

import (
	"encoding/json"
)

type CouponObject struct {
	Id                          int         `json:"id"`
	UniqueId                    string      `json:"uniqid"`
	ShopId                      int         `json:"shop_id"`
	Type                        string      `json:"type"`
	Code                        string      `json:"code"`
	UseType                     string      `json:"use_type"`
	Discount                    int         `json:"discount"`
	Currency                    interface{} `json:"currency"`
	Used                        int         `json:"used"`
	DisabledWithVolumeDiscounts bool        `json:"disabled_with_volume_discounts"`
	AllRecurringBillInvoices    bool        `json:"all_recurring_bill_invoices"`
	MaxUses                     int         `json:"max_uses"`
	ExpireAt                    interface{} `json:"expire_at"`
	ProductsBound               []string    `json:"products_bound"`
	ProductsCount               int         `json:"products_count"`
	CreatedAt                   int64       `json:"created_at"`
	UpdatedAt                   int64       `json:"updated_at"`
	UpdatedBy                   int         `json:"updated_by"`
}

type GetCouponResponseType struct {
	SellixResponseType
	Data *CouponObject
}

func (c *SellixClient) GetCoupon(couponId string) (*GetCouponResponseType, error) {
	body, err := c.createRequest("GET", "/coupons/"+couponId, nil)
	if err != nil {
		return nil, err
	}
	var sellixResponse GetCouponResponseType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

type ListCouponsResponseType struct {
	SellixResponseType
	Data []BlacklistObject
}

func (c *SellixClient) ListCoupons() (*ListCouponsResponseType, error) {
	body, err := c.createRequest("GET", "/coupons", nil)
	if err != nil {
		return nil, err
	}

	var sellixResponse ListCouponsResponseType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

type CouponEntry struct {
	Code                        string   `json:"code"`
	DiscountValue               int      `json:"discount_value"`
	MaxUses                     int      `json:"max_uses"`
	ProductsBound               []string `json:"products_bound"`
	DiscountType                string   `json:"discount_type"`
	DiscountOrderType           string   `json:"discount_order_type"`
	DisabledWithVolumeDiscounts bool     `json:"disabled_with_volume_discounts"`
	AllRecurringBillInvoices    bool     `json:"all_recurring_bill_invoices"`
}

func (c *SellixClient) CreateCoupon(p *CouponEntry) (*PostRequestSellixResponse, error) {
	body, err := c.createRequest("POST", "/coupons", p)
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

func (c *SellixClient) EditCoupon(p *CouponEntry) (*SellixResponseType, error) {
	body, err := c.createRequest("PUT", "/coupons", p)
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

func (c *SellixClient) DeleteCoupon(couponId string) (*SellixResponseType, error) {
	body, err := c.createRequest("DELETE", "/coupons/"+couponId, nil)
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
