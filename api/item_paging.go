package api

import (
	"context"
	"fmt"
)

type ItemPagingReq struct {
	Page            int    `json:"Page"`
	Limit           int    `json:"Limit"`
	BranchID        string `json:"BranchId"`
	CategoryID      string `json:"CategoryId"`
	KeySearch       string `json:"KeySearch"`
	IncludeInactive bool   `json:"includeInactive"`
}

type Item struct {
	ID               string  `json:"Id"`
	Code             string  `json:"Code"`
	ItemType         int     `json:"ItemType"`
	Name             string  `json:"Name"`
	CategoryID       string  `json:"CategoryID"`
	CategoryName     string  `json:"CategoryName"`
	Price            float64 `json:"Price"`
	Inactive         bool    `json:"Inactive"`
	UnitID           string  `json:"UnitID"`
	UnitName         string  `json:"UnitName"`
	Description      string  `json:"Description"`
	IsSeftPrice      bool    `json:"IsSeftPrice"`
	AllowAdjustPrice bool    `json:"AllowAdjustPrice"`
}

func (a *Api) ItemPaging(ctx context.Context, req *ItemPagingReq) (*Response[[]*Item], error) {
	tk, err := a.auth.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := post[responseCommon[[]*Item]](
		ctx,
		fmt.Sprintf(`%s/api/v1/inventoryitems/paging`, a.config.url),
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

	return &Response[[]*Item]{
		Code:    res.Code,
		Data:    res.Data,
		Total:   res.Total,
		Success: res.Success,
	}, nil
}
