package SellixGoSDK

import "encoding/json"

type SubscriptionObject struct {
	Id           string      `json:"id"`
	ShopId       int         `json:"shop_id"`
	ProductId    string      `json:"product_id"`
	Status       string      `json:"status"`
	Gateway      GatewayEnum `json:"gateway"`
	CustomFields struct {
		Username string `json:"username"`
	} `json:"custom_fields"`
	CustomerId                    string      `json:"customer_id"`
	StripeCustomerId              string      `json:"stripe_customer_id"`
	StripeAccount                 string      `json:"stripe_account"`
	StripeSubscriptionId          string      `json:"stripe_subscription_id"`
	CouponId                      interface{} `json:"coupon_id"`
	CurrentPeriodEnd              int         `json:"current_period_end"`
	UpcomingEmail1WeekSent        bool        `json:"upcoming_email_1_week_sent"`
	TrialPeriodEndingEmailSent    bool        `json:"trial_period_ending_email_sent"`
	RenewalInvoiceCreated         bool        `json:"renewal_invoice_created"`
	CreatedAt                     int         `json:"created_at"`
	UpdatedAt                     int         `json:"updated_at"`
	CanceledAt                    interface{} `json:"canceled_at"`
	ProductTitle                  string      `json:"product_title"`
	CustomerName                  string      `json:"customer_name"`
	CustomerSurname               string      `json:"customer_surname"`
	CustomerPhone                 interface{} `json:"customer_phone"`
	CustomerPhoneCountryCode      interface{} `json:"customer_phone_country_code"`
	CustomerCountryCode           interface{} `json:"customer_country_code"`
	CustomerStreetAddress         interface{} `json:"customer_street_address"`
	CustomerAdditionalAddressInfo interface{} `json:"customer_additional_address_info"`
	CustomerCity                  interface{} `json:"customer_city"`
	CustomerPostalCode            interface{} `json:"customer_postal_code"`
	CustomerState                 interface{} `json:"customer_state"`
	CustomerEmail                 string      `json:"customer_email"`
}

type GetSubscriptionResponseType struct {
	SellixResponseType
	Data *QueryObject
}

func (c *SellixClient) GetSubscription(subscriptionId string) (*GetSubscriptionResponseType, error) {
	body, err := c.createRequest("GET", "/subscriptions/"+subscriptionId, nil)
	if err != nil {
		return nil, err
	}
	var sellixResponse GetSubscriptionResponseType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

type ListSubscriptionsResponseType struct {
	SellixResponseType
	Data []QueryObject
}

func (c *SellixClient) ListSubscriptions() (*ListSubscriptionsResponseType, error) {
	body, err := c.createRequest("GET", "/subscriptions", nil)
	if err != nil {
		return nil, err
	}

	var sellixResponse ListSubscriptionsResponseType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

type SubscriptionInfo struct {
	ProductId    string      `json:"product_id"`
	CouponCode   interface{} `json:"coupon_code"`
	CustomFields struct {
		Username string `json:"username"`
	} `json:"custom_fields"`
	CustomerId string      `json:"customer_id"`
	Gateway    interface{} `json:"gateway"`
}

func (c *SellixClient) CreateSubscription(p *SubscriptionInfo) (*SellixResponseType, error) {
	body, err := c.createRequest("POST", "/subscriptions", p)
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

func (c *SellixClient) DeleteSubscription(subscriptionId string) (*SellixResponseType, error) {
	body, err := c.createRequest("DELETE", "/subscriptions/"+subscriptionId, nil)
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
