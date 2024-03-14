package sellix

import "encoding/json"

type CustomerObject struct {
	Id                    string      `json:"id"`
	ShopId                int         `json:"shop_id"`
	Email                 string      `json:"email"`
	Name                  string      `json:"name"`
	Surname               string      `json:"surname"`
	Phone                 int64       `json:"phone"`
	PhoneCountryCode      string      `json:"phone_country_code"`
	CountryCode           string      `json:"country_code"`
	StreetAddress         string      `json:"street_address"`
	AdditionalAddressInfo interface{} `json:"additional_address_info"`
	City                  string      `json:"city"`
	PostalCode            int         `json:"postal_code"`
	State                 string      `json:"state"`
	CreatedAt             int         `json:"created_at"`
}

type GetCustomerResponseType struct {
	SellixResponseType
	Data *GroupObject
}

func (c *SellixClient) GetCustomer(customerId string) (*GetCustomerResponseType, error) {
	body, err := c.createRequest("GET", "/cusomers/"+customerId, nil)
	if err != nil {
		return nil, err
	}
	var sellixResponse GetCustomerResponseType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

type ListCustomersResponseType struct {
	SellixResponseType
	Data []GroupObject
}

func (c *SellixClient) ListCustomers() (*ListCustomersResponseType, error) {
	body, err := c.createRequest("GET", "/customers", nil)
	if err != nil {
		return nil, err
	}

	var sellixResponse ListCustomersResponseType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

type CustomerInfo struct {
	Email                 string      `json:"email"`
	Name                  string      `json:"name"`
	Surname               string      `json:"surname"`
	Phone                 int64       `json:"phone"`
	PhoneCountryCode      string      `json:"phone_country_code"`
	CountryCode           string      `json:"country_code"`
	StreetAddress         string      `json:"street_address"`
	AdditionalAddressInfo interface{} `json:"additional_address_info"`
	City                  string      `json:"city"`
	PostalCode            int         `json:"postal_code"`
	State                 string      `json:"state"`
}

func (c *SellixClient) CreateCustomer(p *CustomerInfo) (*SellixResponseType, error) {
	body, err := c.createRequest("POST", "/customers", p)
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

func (c *SellixClient) EditCustomer(customerId string, p *CustomerInfo) (*SellixResponseType, error) {
	body, err := c.createRequest("PUT", "/customers/"+customerId, p)
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
