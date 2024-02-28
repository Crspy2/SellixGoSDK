package SellixGoSDK

import (
	"encoding/json"
	"fmt"
)

type CheckoutSessionParams struct {
	Title     string `json:"title"`
	ProductID string `json:"product_id"`
	Cart      *struct {
		Products []struct {
			UniqueID     string `json:"uniquid"`
			UnitQuantity uint   `json:"unit_quantity"`
		} `json:"products"`
	} `json:"cart"`
	Gateway          string            `json:"gateway"`
	Gateways         []string          `json:"gateways"`
	PayPalAPM        string            `json:"paypal_apm"`
	CreditCard       bool              `json:"credit_card"`
	LexPaymentMethod string            `json:"lex_payment_method"`
	Value            float64           `json:"value"`
	Currency         string            `json:"currency"`
	Quantity         uint              `json:"quantity"`
	CouponCode       string            `json:"coupon_code"`
	Confirmations    int               `json:"confirmations"`
	Email            string            `json:"email"`
	CustomFields     map[string]string `json:"custom_fields"`
	ProductAddons    map[string]string `json:"product_addons"`
	ProductVariants  map[string]string `json:"product_variants"`
	FraudShield      *struct {
		IP           string `json:"ip"`
		UserAgent    string `json:"user_agent"`
		UserLanguage string `json:"user_language"`
	} `json:"fraud_shield"`
	Webhook    string `json:"webhook"`
	WhiteLabel bool   `json:"white_label"`
	ReturnUrl  string `json:"return_url"`
}

type PaymentNonWhiteLabelResponse struct {
	Status int `json:"status"`
	Data   struct {
		URL      string `json:"url"`
		UniqueID string `json:"uniqid"`
	} `json:"data"`
	Message string      `json:"message"`
	Log     interface{} `json:"log"`
	Error   interface{} `json:"error"`
	Env     string      `json:"env"`
}

type PaymentWhiteLabelResponse struct {
	Status int `json:"status"`
	Data   struct {
		Invoice struct {
			ID                        int               `json:"id"`
			UniqueID                  string            `json:"uniqid"`
			Total                     float64           `json:"total"`
			TotalDisplay              float64           `json:"total_display"`
			Currency                  string            `json:"currency"`
			ExchangeRate              float64           `json:"exchange_rate"`
			CryptoExchangeRate        float64           `json:"crypto_exchange_rate"`
			ShopID                    int               `json:"shop_id"`
			Name                      string            `json:"name"`
			CustomerEmail             string            `json:"customer_email"`
			ProductID                 string            `json:"product_id"`
			ProductType               string            `json:"product_type"`
			ProductPrice              float64           `json:"product_price"`
			FileAttachmentUniqueID    interface{}       `json:"file_attachment_uniqid"`
			Gateway                   string            `json:"gateway"`
			PaypalEmail               interface{}       `json:"paypal_email"`
			PaypalOrderID             interface{}       `json:"paypal_order_id"`
			PaypalPayerEmail          interface{}       `json:"paypal_payer_email"`
			SkrillEmail               interface{}       `json:"skrill_email"`
			SkrillSid                 interface{}       `json:"skrill_sid"`
			SkrillLink                interface{}       `json:"skrill_link"`
			PerfectmoneyID            interface{}       `json:"perfectmoney_id"`
			CryptoAddress             string            `json:"crypto_address"`
			CryptoAmount              float64           `json:"crypto_amount"`
			CryptoReceived            int               `json:"crypto_received"`
			CryptoURI                 string            `json:"crypto_uri"`
			CryptoConfirmationsNeeded int               `json:"crypto_confirmations_needed"`
			Country                   string            `json:"country"`
			Location                  string            `json:"location"`
			IP                        string            `json:"ip"`
			IsVPNOrProxy              bool              `json:"is_vpn_or_proxy"`
			UserAgent                 string            `json:"user_agent"`
			Quantity                  int               `json:"quantity"`
			CouponID                  interface{}       `json:"coupon_id"`
			CustomFields              map[string]string `json:"custom_fields"`
			DeveloperInvoice          bool              `json:"developer_invoice"`
			DeveloperTitle            string            `json:"developer_title"`
			DeveloperWebhook          string            `json:"developer_webhook"`
			DeveloperReturnURL        string            `json:"developer_return_url"`
			Status                    string            `json:"status"`
			Discount                  float64           `json:"discount"`
			FeeFixed                  float64           `json:"fee_fixed"`
			FeePercentage             float64           `json:"fee_percentage"`
			DayValue                  int               `json:"day_value"`
			Day                       string            `json:"day"`
			Month                     string            `json:"month"`
			Year                      int               `json:"year"`
			CreatedAt                 int               `json:"created_at"`
			UpdatedAt                 int               `json:"updated_at"`
			UpdatedBy                 int               `json:"updated_by"`
			Serials                   []interface{}     `json:"serials"`
			File                      interface{}       `json:"file"`
			Webhooks                  []interface{}     `json:"webhooks"`
			CryptoPayout              bool              `json:"crypto_payout"`
			CryptoPayoutTransaction   interface{}       `json:"crypto_payout_transaction"`
			CryptoTransactions        []interface{}     `json:"crypto_transactions"`
			Product                   struct {
				Title        string  `json:"title"`
				PriceDisplay float64 `json:"price_display"`
				Currency     string  `json:"currency"`
			} `json:"product"`
			TotalConversions map[string]float64 `json:"total_conversions"`
			Theme            string             `json:"theme"`
		}
	} `json:"data"`
}

// CreatePaymentNonWhitelabel creates a payment session with CheckoutSessionParams.whitelabel=false
func (c *SellixClient) CreatePaymentNonWhitelabel(payment *CheckoutSessionParams) (*PaymentNonWhiteLabelResponse, error) {
	payment.WhiteLabel = false

	body, err := c.createRequest("POST", "/payments", payment)
	if err != nil {
		return nil, err
	}

	var sellixResponse PaymentNonWhiteLabelResponse
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return &sellixResponse, nil
}

func (c *SellixClient) CreatePaymentWhitelabel(payment *CheckoutSessionParams) (*PaymentWhiteLabelResponse, error) {
	payment.WhiteLabel = true

	body, err := c.createRequest("POST", "/payments", payment)
	if err != nil {
		return nil, err
	}

	var sellixResponse PaymentWhiteLabelResponse
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

func (c *SellixClient) CompletePayment(uniqueId string) (*SellixReturnType, error) {
	body, err := c.createRequest("PUT", "/payments/"+uniqueId, nil)
	if err != nil {
		return nil, err
	}

	var sellixResponse SellixReturnType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

func (c *SellixClient) DeletePayment(uniqueId string) (*SellixReturnType, error) {
	body, err := c.createRequest("DELETE", "/payments/"+uniqueId, nil)
	if err != nil {
		return nil, err
	}

	var sellixResponse SellixReturnType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}
