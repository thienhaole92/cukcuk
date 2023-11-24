package cukcuk

import (
	"context"
	"fmt"
)

type SAInvoiceDetails struct {
	RefDetailID                       string  `json:"RefDetailID"`
	RefDetailType                     int     `json:"RefDetailType"`
	RefID                             string  `json:"RefID"`
	ItemID                            string  `json:"ItemID"`
	ItemName                          string  `json:"ItemName"`
	Quantity                          float64 `json:"Quantity"`
	UnitPrice                         float64 `json:"UnitPrice"`
	UnitID                            string  `json:"UnitID"`
	UnitName                          string  `json:"UnitName"`
	Amount                            float64 `json:"Amount"`
	DiscountRate                      float64 `json:"DiscountRate"`
	SortOrder                         int     `json:"SortOrder"`
	InventoryItemType                 int     `json:"InventoryItemType"`
	HaveAddition                      bool    `json:"HaveAddition"`
	IsSeftPrice                       bool    `json:"IsSeftPrice"`
	PromotionRate                     float64 `json:"PromotionRate"`
	PromotionType                     int     `json:"PromotionType"`
	PromotionName                     string  `json:"PromotionName"`
	OrderDetailID                     string  `json:"OrderDetailID"`
	IsSelected                        bool    `json:"IsSelected"`
	SAInvoicePromotionAmount          float64 `json:"SAInvoicePromotionAmount"`
	ItemCode                          string  `json:"ItemCode"`
	PromotionAmount                   float64 `json:"PromotionAmount"`
	AllocationAmount                  float64 `json:"AllocationAmount"`
	PreTaxAmount                      float64 `json:"PreTaxAmount"`
	AllocationDeliveryPromotionAmount float64 `json:"AllocationDeliveryPromotionAmount"`
}

type SAInvoicePayments struct {
	SAInvoicePaymentID string  `json:"SAInvoicePaymentID"`
	RefID              string  `json:"RefID"`
	PaymentType        int     `json:"PaymentType"`
	Amount             float64 `json:"Amount"`
	PaymentName        string  `json:"PaymentName"`
	FoodAmount         float64 `json:"FoodAmount"`
	DrinkAmount        float64 `json:"DrinkAmount"`
	ApplyVoucherType   int     `json:"ApplyVoucherType"`
	VoucherAllAmount   float64 `json:"VoucherAllAmount"`
	VoucherFoodAmount  float64 `json:"VoucherFoodAmount"`
	VoucherDrinkAmount float64 `json:"VoucherDrinkAmount"`
	ExchangeRate       float64 `json:"ExchangeRate"`
	ExchangeAmount     float64 `json:"ExchangeAmount"`
}

type SAVATInfo struct {
	Vatid                             string `json:"VATID"`
	RefID                             string `json:"RefID"`
	ReceiverEIvoiceName               string `json:"ReceiverEIvoiceName"`
	Tel                               string `json:"Tel"`
	CompanyName                       string `json:"CompanyName"`
	CompanyAddress                    string `json:"CompanyAddress"`
	TaxCode                           string `json:"TaxCode"`
	Email                             string `json:"Email"`
	Status                            bool   `json:"Status"`
	StatusReleaseEInvoice             int    `json:"StatusReleaseEInvoice"`
	EInvoiceNumber                    string `json:"EInvoiceNumber"`
	StatusSendEmail                   int    `json:"StatusSendEmail"`
	TransactionID                     string `json:"TransactionID"`
	SellerTaxCode                     string `json:"SellerTaxCode"`
	TemplateCode                      string `json:"TemplateCode"`
	InvoiceSeries                     string `json:"InvoiceSeries"`
	RefDateReleaseEInvoice            string `json:"RefDateReleaseEInvoice"`
	StatusSendToTax                   int    `json:"StatusSendToTax"`
	AccountObjectIdentificationNumber string `json:"AccountObjectIdentificationNumber"`
	IsCalculatingMachinePublishing    bool   `json:"IsCalculatingMachinePublishing"`
	Explanation                       string `json:"Explanation"`
}

type SAInvoice struct {
	RefID                   string               `json:"RefId"`
	RefType                 int                  `json:"RefType"`
	RefNo                   string               `json:"RefNo"`
	RefDate                 string               `json:"RefDate"`
	BranchID                string               `json:"BranchId"`
	OrderID                 string               `json:"OrderId"`
	OrderType               int                  `json:"OrderType"`
	EmployeeID              string               `json:"EmployeeId"`
	EmployeeName            string               `json:"EmployeeName"`
	WaiterEmployeeName      string               `json:"WaiterEmployeeName"`
	TableName               string               `json:"TableName"`
	Description             string               `json:"Description"`
	DepositAmount           float64              `json:"DepositAmount"`
	Amount                  float64              `json:"Amount"`
	DeliveryAmount          float64              `json:"DeliveryAmount"`
	ServiceRate             float64              `json:"ServiceRate"`
	ServiceAmount           float64              `json:"ServiceAmount"`
	VATRate                 float64              `json:"VATRate"`
	VATAmount               float64              `json:"VATAmount"`
	DiscountAmount          float64              `json:"DiscountAmount"`
	PromotionRate           float64              `json:"PromotionRate"`
	PromotionAmount         float64              `json:"PromotionAmount"`
	PromotionItemsAmount    float64              `json:"PromotionItemsAmount"`
	ReceiveAmount           float64              `json:"ReceiveAmount"`
	ReturnAmount            float64              `json:"ReturnAmount"`
	TotalAmount             float64              `json:"TotalAmount"`
	SaleAmount              float64              `json:"SaleAmount"`
	TotalItemAmount         float64              `json:"TotalItemAmount"`
	TotalItemAmountAfterTax float64              `json:"TotalItemAmountAfterTax"`
	TipAmount               float64              `json:"TipAmount"`
	ServiceTaxRate          float64              `json:"ServiceTaxRate"`
	DeliveryTaxRate         float64              `json:"DeliveryTaxRate"`
	PaymentStatus           int                  `json:"PaymentStatus"`
	AvailablePoint          int                  `json:"AvailablePoint"`
	UsedPoint               int                  `json:"UsedPoint"`
	AddPoint                int                  `json:"AddPoint"`
	SAInvoiceDetails        []*SAInvoiceDetails  `json:"SAInvoiceDetails"`
	SAInvoicePayments       []*SAInvoicePayments `json:"SAInvoicePayments"`
	SAVATInfo               *SAVATInfo           `json:"SAVATInfo"`
}

func (a *Api) GetSAInvoice(ctx context.Context, refId string) (*Response[*SAInvoice], error) {
	tk, err := a.auth.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := get[responseCommon[*SAInvoice]](
		ctx,
		fmt.Sprintf(`%s/api/v1/sainvoices/%s`, a.config.url, refId),
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

	return &Response[*SAInvoice]{
		Code:    res.Code,
		Data:    res.Data,
		Total:   res.Total,
		Success: res.Success,
	}, nil
}
