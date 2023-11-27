package api

import (
	"context"
	"fmt"
)

type CreateCustomerReq struct {
	Name string `json:"Name"`
	Tel  string `json:"Tel"`
}

func (a *Api) CreateCustomer(ctx context.Context, req *CreateCustomerReq) (*Response[[]*Customer], error) {
	tk, err := a.auth.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := post[responseCommon[[]*Customer]](
		ctx,
		fmt.Sprintf(`%s/api/v1/customers`, a.config.url),
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
