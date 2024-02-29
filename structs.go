package SellixGoSDK

import "net/http"

type SellixClient struct {
	Http      *http.Client
	Request   *http.Request
	ApiKey    string
	storeName *string
	baseUrl   string
}

type SellixResponseType struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Log     interface{} `json:"log"`
	Error   interface{} `json:"error"`
	Env     string      `json:"env"`
}

type PostRequestSellixResponse struct {
	SellixResponseType
	Data *struct {
		UniqueId string `json:"uniqid"`
	}
}

type GatewayEnum string

const (
	GatewayPayPal       GatewayEnum = "PAYPAL"
	GatewayStripe       GatewayEnum = "STRIPE"
	GatewaySkrill       GatewayEnum = "SKRILL"
	GatewayPerfectMoney GatewayEnum = "PERFECT_MONEY"
	GatewayCashApp      GatewayEnum = "CASH_APP"
	GatewaySolana       GatewayEnum = "SOLANA"
	GatewayRipple       GatewayEnum = "RIPPLE"
	GatewayBitcoin      GatewayEnum = "BITCOIN"
	GatewayEthereum     GatewayEnum = "ETHEREUM"
	GatewayLitecoin     GatewayEnum = "LITECOIN"
	GatewayTron         GatewayEnum = "TRON"
	GatewayBitcoinCash  GatewayEnum = "BITCOIN_CASH"
	GatewayBinanceCoin  GatewayEnum = "BINANCE_COIN"
	GatewayMonero       GatewayEnum = "MONERO"
	GatewayBitcoinLN    GatewayEnum = "BITCOIN_LN"
	GatewayConcordium   GatewayEnum = "CONCORDIUM"
	GatewayDogecoin     GatewayEnum = "DOGECOIN"
	GatewayAPE          GatewayEnum = "APE"
	GatewaySHIB         GatewayEnum = "SHIB"
	GatewayPEPE         GatewayEnum = "PEPE"
	GatewayDAI          GatewayEnum = "DAI"
	GatewayNano         GatewayEnum = "NANO"
	GatewayPolygon      GatewayEnum = "POLYGON"
	GatewayPLZTRC20     GatewayEnum = "PLZ:TRC20"
	GatewayPLZBEP20     GatewayEnum = "PLZ:BEP20"
	GatewayUSDCMATIC    GatewayEnum = "USDC:MATIC"
	GatewayUSDTMATIC    GatewayEnum = "USDT:MATIC"
	GatewayUSDCERC20    GatewayEnum = "USDC:ERC20"
	GatewayUSDCBEP20    GatewayEnum = "USDC:BEP20"
	GatewayUSDTERC20    GatewayEnum = "USDT:ERC20"
	GatewayUSDTTRC20    GatewayEnum = "USDT:TRC20"
	GatewayUSDTBEP20    GatewayEnum = "USDT:BEP20"
	GatewayCronos       GatewayEnum = "CRONOS"
)

type PaypalAPMEnum string

const (
	PaypalAPMBancontact PaypalAPMEnum = "bancontact"
	PaypalAPMEps        PaypalAPMEnum = "eps"
	PaypalAPMTrustly    PaypalAPMEnum = "trustly"
	PaypalAPMMercado    PaypalAPMEnum = "mercado"
	PaypalAPMPaylater   PaypalAPMEnum = "paylater"
	PaypalAPMSepa       PaypalAPMEnum = "sepa"
	PaypalAPMVenmo      PaypalAPMEnum = "venmo"
	PaypalAPMBlik       PaypalAPMEnum = "blik"
	PaypalAPMGiropay    PaypalAPMEnum = "giropay"
	PaypalAPMIdeal      PaypalAPMEnum = "ideal"
	PaypalAPMMybank     PaypalAPMEnum = "mybank"
	PaypalAPMSofort     PaypalAPMEnum = "sofort"
	PaypalAPMPrzelewy24 PaypalAPMEnum = "przelewy24"
	PaypalAPMCredit     PaypalAPMEnum = "credit"
)

type StripeAPMEnum string

const (
	StripeAPMAfterpayClearpay StripeAPMEnum = "afterpay_clearpay"
	StripeAPMAlipay           StripeAPMEnum = "alipay"
	StripeAPMBancontact       StripeAPMEnum = "bancontact"
	StripeAPMAUBecsDebit      StripeAPMEnum = "au_becs_debit"
	StripeAPMBoleto           StripeAPMEnum = "boleto"
	StripeAPMEps              StripeAPMEnum = "eps"
	StripeAPMFpx              StripeAPMEnum = "fpx"
	StripeAPMGiropay          StripeAPMEnum = "giropay"
	StripeAPMGrabpay          StripeAPMEnum = "grabpay"
	StripeAPMIdeal            StripeAPMEnum = "ideal"
	StripeAPMKlarna           StripeAPMEnum = "klarna"
	StripeAPMOxxo             StripeAPMEnum = "oxxo"
	StripeAPMP24              StripeAPMEnum = "p24"
	StripeAPMSofort           StripeAPMEnum = "sofort"
	StripeAPMWechatPay        StripeAPMEnum = "wechat_pay"
)
