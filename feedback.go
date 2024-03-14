package sellix

import (
	"encoding/json"
)

type FeedbackObject struct {
	Id                  int         `json:"id"`
	UniqueId            string      `json:"uniqid"`
	ProductId           string      `json:"product_id"`
	InvoiceId           string      `json:"invoice_id"`
	Blocked             bool        `json:"blocked"`
	Appealed            bool        `json:"appealed"`
	ShopId              int         `json:"shop_id"`
	Message             string      `json:"message"`
	Reply               string      `json:"reply"`
	Score               int         `json:"score"`
	ProductTitle        string      `json:"product_title"`
	ProductImageName    interface{} `json:"product_image_name"`
	ProductImageStorage interface{} `json:"product_image_storage"`
	CloudflareImageId   string      `json:"cloudflare_image_id"`
	Invoice             struct {
		Id                        int         `json:"id"`
		UniqueId                  string      `json:"uniqid"`
		RecurringBillingId        interface{} `json:"recurring_billing_id"`
		Total                     float64     `json:"total"`
		TotalDisplay              int         `json:"total_display"`
		ExchangeRate              float64     `json:"exchange_rate"`
		CryptoExchangeRate        float64     `json:"crypto_exchange_rate"`
		Currency                  string      `json:"currency"`
		ShopId                    int         `json:"shop_id"`
		ShopImageName             interface{} `json:"shop_image_name"`
		ShopImageStorage          interface{} `json:"shop_image_storage"`
		CloudflareImageId         string      `json:"cloudflare_image_id"`
		Name                      string      `json:"name"`
		Type                      string      `json:"type"`
		CustomerEmail             string      `json:"customer_email"`
		PaypalEmailDelivery       bool        `json:"paypal_email_delivery"`
		ProductId                 string      `json:"product_id"`
		ProductTitle              string      `json:"product_title"`
		ProductType               string      `json:"product_type"`
		Subtype                   interface{} `json:"subtype"`
		SubscriptionId            interface{} `json:"subscription_id"`
		SubscriptionTime          interface{} `json:"subscription_time"`
		Gateway                   string      `json:"gateway"`
		PaypalApm                 interface{} `json:"paypal_apm"`
		PaypalEmail               interface{} `json:"paypal_email"`
		PaypalOrderId             interface{} `json:"paypal_order_id"`
		PaypalFee                 interface{} `json:"paypal_fee"`
		PaypalPayerEmail          interface{} `json:"paypal_payer_email"`
		PaypalSubscriptionId      interface{} `json:"paypal_subscription_id"`
		PaypalSubscriptionLink    interface{} `json:"paypal_subscription_link"`
		LexOrderId                interface{} `json:"lex_order_id"`
		LexPaymentMethod          interface{} `json:"lex_payment_method"`
		StripeClientSecret        interface{} `json:"stripe_client_secret"`
		StripePriceId             interface{} `json:"stripe_price_id"`
		SkrillEmail               interface{} `json:"skrill_email"`
		SkrillSid                 interface{} `json:"skrill_sid"`
		SkrillLink                interface{} `json:"skrill_link"`
		PerfectmoneyId            interface{} `json:"perfectmoney_id"`
		CryptoAddress             string      `json:"crypto_address"`
		CryptoAmount              float64     `json:"crypto_amount"`
		CryptoReceived            float64     `json:"crypto_received"`
		CryptoUri                 string      `json:"crypto_uri"`
		CryptoConfirmationsNeeded int         `json:"crypto_confirmations_needed"`
		CryptoScheduledPayout     bool        `json:"crypto_scheduled_payout"`
		CryptoPayout              bool        `json:"crypto_payout"`
		FeeBilled                 bool        `json:"fee_billed"`
		BillInfo                  interface{} `json:"bill_info"`
		CashappQrcode             interface{} `json:"cashapp_qrcode"`
		CashappCashtag            interface{} `json:"cashapp_cashtag"`
		CashappNote               interface{} `json:"cashapp_note"`
		Country                   string      `json:"country"`
		Location                  string      `json:"location"`
		Ip                        string      `json:"ip"`
		IsVpnOrProxy              bool        `json:"is_vpn_or_proxy"`
		UserAgent                 string      `json:"user_agent"`
		Quantity                  int         `json:"quantity"`
		CouponId                  interface{} `json:"coupon_id"`
		CustomFields              struct {
			Username string `json:"username"`
		} `json:"custom_fields"`
		DeveloperInvoice   bool        `json:"developer_invoice"`
		DeveloperTitle     interface{} `json:"developer_title"`
		DeveloperWebhook   interface{} `json:"developer_webhook"`
		DeveloperReturnUrl interface{} `json:"developer_return_url"`
		Status             string      `json:"status"`
		StatusDetails      interface{} `json:"status_details"`
		VoidDetails        interface{} `json:"void_details"`
		Discount           int         `json:"discount"`
		FeePercentage      int         `json:"fee_percentage"`
		IpInfo             struct {
			Success         bool   `json:"success"`
			Message         string `json:"message"`
			FraudScore      int    `json:"fraud_score"`
			CountryCode     string `json:"country_code"`
			Region          string `json:"region"`
			City            string `json:"city"`
			ISP             string `json:"ISP"`
			ASN             int    `json:"ASN"`
			OperatingSystem string `json:"operating_system"`
			Browser         string `json:"browser"`
			Organization    string `json:"organization"`
			IsCrawler       bool   `json:"is_crawler"`
			Timezone        string `json:"timezone"`
			Mobile          bool   `json:"mobile"`
			Host            string `json:"host"`
			Proxy           bool   `json:"proxy"`
			Vpn             bool   `json:"vpn"`
			Tor             bool   `json:"tor"`
			ActiveVpn       bool   `json:"active_vpn"`
			ActiveTor       bool   `json:"active_tor"`
			DeviceBrand     string `json:"device_brand"`
			DeviceModel     string `json:"device_model"`
			RecentAbuse     bool   `json:"recent_abuse"`
			BotStatus       bool   `json:"bot_status"`
			ConnectionType  string `json:"connection_type"`
			AbuseVelocity   string `json:"abuse_velocity"`
			ZipCode         string `json:"zip_code"`
			Latitude        int    `json:"latitude"`
			Longitude       int    `json:"longitude"`
			RequestId       string `json:"request_id"`
		} `json:"ip_info"`
		Serials []string `json:"serials"`
		File    struct {
			Id                int         `json:"id"`
			UniqueId          string      `json:"uniqid"`
			CloudflareImageId string      `json:"cloudflare_image_id"`
			Storage           interface{} `json:"storage"`
			Name              string      `json:"name"`
			OriginalName      string      `json:"original_name"`
			Extension         string      `json:"extension"`
			ShopId            int         `json:"shop_id"`
			Size              int         `json:"size"`
			CreatedAt         int         `json:"created_at"`
		} `json:"file"`
		ServiceText     interface{} `json:"service_text"`
		DynamicResponse interface{} `json:"dynamic_response"`
		Webhooks        []struct {
			UniqueId     string `json:"uniqid"`
			Url          string `json:"url"`
			Event        string `json:"event"`
			Retries      int    `json:"retries"`
			ResponseCode int    `json:"response_code"`
			CreatedAt    int    `json:"created_at"`
			Payload      struct {
				Sample string `json:"sample"`
			} `json:"payload"`
			Response string `json:"response"`
		} `json:"webhooks"`
		CryptoPayoutTransaction struct {
			ToAddress    string  `json:"to_address"`
			FromAddress  string  `json:"from_address"`
			CryptoAmount float64 `json:"crypto_amount"`
			Hash         string  `json:"hash"`
			CreatedAt    int     `json:"created_at"`
		} `json:"crypto_payout_transaction"`
		PaypalDispute struct {
			Id        string      `json:"id"`
			InvoiceId string      `json:"invoice_id"`
			ShopId    int         `json:"shop_id"`
			Reason    string      `json:"reason"`
			Status    string      `json:"status"`
			Outcome   interface{} `json:"outcome"`
			Messages  []struct {
				PostedBy  string `json:"posted_by"`
				Content   string `json:"content"`
				CreatedAt int    `json:"created_at"`
			} `json:"messages"`
			LifeCycleStage        string `json:"life_cycle_stage"`
			SellerResponseDueDate int    `json:"seller_response_due_date"`
			CreatedAt             int64  `json:"created_at"`
			UpdatedAt             int64  `json:"updated_at"`
		} `json:"paypal_dispute"`
		StatusHistory []struct {
			Id        int    `json:"id"`
			InvoiceId string `json:"invoice_id"`
			Status    string `json:"status"`
			Details   string `json:"details"`
			CreatedAt int    `json:"created_at"`
		} `json:"status_history"`
		CryptoTransactions []struct {
			CryptoAmount  float64 `json:"crypto_amount"`
			Hash          string  `json:"hash"`
			Confirmations int     `json:"confirmations"`
			CreatedAt     int     `json:"created_at"`
			UpdatedAt     int     `json:"updated_at"`
		} `json:"crypto_transactions"`
		GatewaysAvailable            []string `json:"gateways_available"`
		ShopPaypalCreditCard         bool     `json:"shop_paypal_credit_card"`
		ShopForcePaypalEmailDelivery bool     `json:"shop_force_paypal_email_delivery"`
		Product                      struct {
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
		} `json:"product"`
		DayValue  int    `json:"day_value"`
		Day       string `json:"day"`
		Month     string `json:"month"`
		Year      int    `json:"year"`
		CreatedAt int64  `json:"created_at"`
		UpdatedAt int64  `json:"updated_at"`
		UpdatedBy int    `json:"updated_by"`
	} `json:"invoice"`
	Product struct {
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
	} `json:"product"`
}

type GetFeedbackResponseType struct {
	SellixResponseType
	Data *CouponObject
}

func (c *SellixClient) GetFeedback(feedbackId string) (*GetCouponResponseType, error) {
	body, err := c.createRequest("GET", "/feedback/"+feedbackId, nil)
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

type ListFeedbackResponseType struct {
	SellixResponseType
	Data []BlacklistObject
}

func (c *SellixClient) ListFeedback() (*ListFeedbackResponseType, error) {
	body, err := c.createRequest("GET", "/feedback", nil)
	if err != nil {
		return nil, err
	}

	var sellixResponse ListFeedbackResponseType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

func (c *SellixClient) ReplyFeedback(feedbackId, reply string) (*PostRequestSellixResponse, error) {
	type FeedbackEntry struct {
		Reply string `json:"reply"`
	}
	p := &FeedbackEntry{
		reply,
	}

	body, err := c.createRequest("POST", "/feedback/reply/"+feedbackId, p)
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
