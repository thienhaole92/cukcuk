package cukcuk

import (
	"context"
	"fmt"
)

type OrderDetails struct {
	ID        string  `json:"Id"`
	ItemID    string  `json:"ItemId"`
	ItemName  string  `json:"ItemName"`
	UnitID    string  `json:"UnitId"`
	UnitName  string  `json:"UnitName"`
	Quantity  float64 `json:"Quantity"`
	Status    int     `json:"Status"`
	Price     float64 `json:"Price"`
	Amount    float64 `json:"Amount"`
	SortOrder int     `json:"SortOrder"`
}

type Order struct {
	ID                 string          `json:"Id"`
	Type               int             `json:"Type"`
	No                 string          `json:"No"`
	BranchID           string          `json:"BranchId"`
	Status             int             `json:"Status"`
	Date               string          `json:"Date"`
	CustomerID         string          `json:"CustomerId"`
	CustomerName       string          `json:"CustomerName"`
	CustomerTel        string          `json:"CustomerTel"`
	EmployeeID         string          `json:"EmployeeId"`
	ShippingAddress    string          `json:"ShippingAddress"`
	DeliveryAmount     float64         `json:"DeliveryAmount"`
	DepositAmount      float64         `json:"DepositAmount"`
	RequestDescription string          `json:"RequestDescription"`
	TotalAmount        float64         `json:"TotalAmount"`
	TableName          string          `json:"TableName"`
	OrderDetails       []*OrderDetails `json:"OrderDetails"`
}

func (a *Api) GetOrder(ctx context.Context, orderId string) (*Response[*Order], error) {
	tk, err := a.auth.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := get[responseCommon[*Order]](
		ctx,
		fmt.Sprintf(`%s/api/v1/orders/%s`, a.config.url, orderId),
		map[string]string{
			"Authorization": fmt.Sprintf(`Bearer %s`, tk),
			"CompanyCode":   a.config.companyCode,
		},
	)

	if err != nil {
		return nil, err
	}

	if !res.Success {
		return nil, fmt.Errorf(res.ErrorMessage)
	}

	return &Response[*Order]{
		Code:    res.Code,
		Data:    res.Data,
		Total:   res.Total,
		Success: res.Success,
	}, nil
}
