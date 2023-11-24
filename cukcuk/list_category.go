package cukcuk

import (
	"context"
	"fmt"
)

type Category struct {
	ID          string `json:"Id"`
	Code        string `json:"Code"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	IsLeaf      bool   `json:"IsLeaf"`
	Grade       int    `json:"Grade"`
	Inactive    bool   `json:"Inactive"`
}

func (a *Api) ListCategory(ctx context.Context, includeInactive bool) (*Response[[]*Category], error) {
	tk, err := a.auth.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := get[responseCommon[[]*Category]](
		ctx,
		fmt.Sprintf(`%s/api/v1/categories/list?includeInactive=%t`, a.config.url, includeInactive),
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

	return &Response[[]*Category]{
		Code:    res.Code,
		Data:    res.Data,
		Total:   res.Total,
		Success: res.Success,
	}, nil
}
