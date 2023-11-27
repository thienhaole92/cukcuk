package api

import (
	"context"
	"fmt"
)

type SAInvoicePagingReq struct {
	Page         int    `json:"Page"`
	Limit        int    `json:"Limit"`
	BranchID     string `json:"BranchId"`
	HaveCustomer bool   `json:"HaveCustomer"`
}

type SAInvoiceDef struct {
	RefID                   string  `json:"RefId"`
	RefType                 int     `json:"RefType"`
	RefNo                   string  `json:"RefNo"`
	RefDate                 string  `json:"RefDate"`
	BranchID                string  `json:"BranchId"`
	OrderType               int     `json:"OrderType"`
	CustomerID              string  `json:"CustomerId"`
	CustomerName            string  `json:"CustomerName"`
	EmployeeID              string  `json:"EmployeeId"`
	EmployeeName            string  `json:"EmployeeName"`
	Description             string  `json:"Description"`
	DepositAmount           float64 `json:"DepositAmount"`
	Amount                  float64 `json:"Amount"`
	DeliveryAmount          float64 `json:"DeliveryAmount"`
	ServiceRate             float64 `json:"ServiceRate"`
	ServiceAmount           float64 `json:"ServiceAmount"`
	VATRate                 float64 `json:"VATRate"`
	VATAmount               float64 `json:"VATAmount"`
	DiscountAmount          float64 `json:"DiscountAmount"`
	PromotionRate           float64 `json:"PromotionRate"`
	PromotionAmount         float64 `json:"PromotionAmount"`
	PromotionItemsAmount    float64 `json:"PromotionItemsAmount"`
	ReceiveAmount           float64 `json:"ReceiveAmount"`
	ReturnAmount            float64 `json:"ReturnAmount"`
	TotalAmount             float64 `json:"TotalAmount"`
	SaleAmount              float64 `json:"SaleAmount"`
	TotalItemAmount         float64 `json:"TotalItemAmount"`
	TotalItemAmountAfterTax float64 `json:"TotalItemAmountAfterTax"`
	TipAmount               float64 `json:"TipAmount"`
	ServiceTaxRate          float64 `json:"ServiceTaxRate"`
	DeliveryTaxRate         float64 `json:"DeliveryTaxRate"`
	PaymentStatus           int     `json:"PaymentStatus"`
	AvailablePoint          int     `json:"AvailablePoint"`
	UsedPoint               int     `json:"UsedPoint"`
	AddPoint                int     `json:"AddPoint"`
}

func (a *Api) SAInvoicePaging(ctx context.Context, req *SAInvoicePagingReq) (*Response[[]*SAInvoiceDef], error) {
	tk, err := a.auth.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := post[responseCommon[[]*SAInvoiceDef]](
		ctx,
		fmt.Sprintf(`%s/api/v1/sainvoices/paging`, a.config.url),
		map[string]string{
			"Authorization": fmt.Sprintf(`Bearer %s`, tk),
			"CompanyCode":   a.config.companyCode,
		},
		req,
	)

	if err != nil {
		return nil, err
	}

	if !res.Success {
		return nil, fmt.Errorf(res.ErrorMessage)
	}

	return &Response[[]*SAInvoiceDef]{
		Code:    res.Code,
		Data:    res.Data,
		Total:   res.Total,
		Success: res.Success,
	}, nil
}
