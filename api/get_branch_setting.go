package api

import (
	"context"
	"fmt"
)

type BranchSetting struct {
	ID                              string  `json:"Id"`
	HasVATRate                      bool    `json:"HasVATRate"`
	VATRate                         float64 `json:"VATRate"`
	VATForDelivery                  bool    `json:"VATForDelivery"`
	VATForTakeAway                  bool    `json:"VATForTakeAway"`
	VATForShip                      bool    `json:"VATForShip"`
	CalTaxForService                bool    `json:"CalTaxForService"`
	HasCalcService                  bool    `json:"HasCalcService"`
	CalcServiceDelivery             bool    `json:"CalcServiceDelivery"`
	CalcServiceTakeAway             bool    `json:"CalcServiceTakeAway"`
	IsCalcServiceAmountNotPromotion bool    `json:"IsCalcServiceAmountNotPromotion"`
	HasCalcServiceWhenRequire       bool    `json:"HasCalcServiceWhenRequire"`
	HasServiceRate                  bool    `json:"HasServiceRate"`
	ServiceRate                     float64 `json:"ServiceRate"`
	HasAmountService                bool    `json:"HasAmountService"`
	AmountService                   float64 `json:"AmountService"`
	Code                            string  `json:"Code"`
	Name                            string  `json:"Name"`
	IsBaseDepot                     bool    `json:"IsBaseDepot"`
	Inactive                        bool    `json:"Inactive"`
}

func (a *Api) GetBranchSetting(ctx context.Context, branchId string) (*Response[*BranchSetting], error) {
	tk, err := a.auth.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := get[responseCommon[*BranchSetting]](
		ctx,
		fmt.Sprintf(`%s/api/v1/branchs/setting/%s`, a.config.url, branchId),
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

	return &Response[*BranchSetting]{
		Code:    res.Code,
		Data:    res.Data,
		Total:   res.Total,
		Success: res.Success,
	}, nil
}
