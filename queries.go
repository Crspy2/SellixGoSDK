package sellix

import "encoding/json"

type QueryObject struct {
	Id            int         `json:"id"`
	UniqueId      string      `json:"uniqid"`
	ShopId        int         `json:"shop_id"`
	InvoiceId     interface{} `json:"invoice_id"`
	CustomerEmail string      `json:"customer_email"`
	Title         string      `json:"title"`
	Status        string      `json:"status"`
	Messages      []struct {
		Role      string `json:"role"`
		Message   string `json:"message"`
		CreatedAt int64  `json:"created_at"`
	} `json:"messages"`
	DayValue  int    `json:"day_value"`
	Day       string `json:"day"`
	Month     string `json:"month"`
	Year      int    `json:"year"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	UpdatedBy int    `json:"updated_by"`
}

type GetQueryResponseType struct {
	SellixResponseType
	Data *QueryObject
}

func (c *SellixClient) GetQuery(queryId string) (*GetQueryResponseType, error) {
	body, err := c.createRequest("GET", "/queries/"+queryId, nil)
	if err != nil {
		return nil, err
	}
	var sellixResponse GetQueryResponseType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

type ListQueryResponseType struct {
	SellixResponseType
	Data []QueryObject
}

func (c *SellixClient) ListQueries() (*ListQueryResponseType, error) {
	body, err := c.createRequest("GET", "/queries", nil)
	if err != nil {
		return nil, err
	}

	var sellixResponse ListQueryResponseType
	err = json.Unmarshal(body, &sellixResponse)
	if err != nil {
		return nil, err
	}

	return &sellixResponse, nil
}

type QueryReplyInfo struct {
	Reply string `json:"reply"`
}

func (c *SellixClient) QueryReply(p *QueryReplyInfo) (*SellixResponseType, error) {
	body, err := c.createRequest("POST", "/queries/reply/", p)
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

func (c *SellixClient) CloseQuery(queryId string) (*SellixResponseType, error) {
	body, err := c.createRequest("POST", "/queries/close/"+queryId, nil)
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

func (c *SellixClient) ReopenQuery(queryId string) (*SellixResponseType, error) {
	body, err := c.createRequest("POST", "/queries/reopen/"+queryId, nil)
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
