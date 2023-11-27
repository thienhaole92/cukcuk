package api

import (
	"context"
	"fmt"
)

type CustomerPagingReq struct {
	Page            int    `json:"Page"`
	Limit           int    `json:"Limit"`
	IncludeInactive bool   `json:"IncludeInactive"`
	LastSyncDate    string `json:"LastSyncDate"`
}

type Customer struct {
	ID          string `json:"Id"`
	Code        string `json:"Code"`
	Name        string `json:"Name"`
	Tel         string `json:"Tel"`
	Birthday    string `json:"Birthday"`
	Address     string `json:"Address"`
	Description string `json:"Description"`
	Email       string `json:"Email"`
	Inactive    bool   `json:"Inactive"`
}

func (a *Api) CustomerPaging(ctx context.Context, req *CustomerPagingReq) (*Response[[]*Customer], error) {
	tk, err := a.auth.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := post[responseCommon[[]*Customer]](
		ctx,
		fmt.Sprintf(`%s/api/v1/customers/paging`, a.config.url),
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

	return &Response[[]*Customer]{
		Code:    res.Code,
		Data:    res.Data,
		Total:   res.Total,
		Success: res.Success,
	}, nil
}
