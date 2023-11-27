package api

import (
	"context"
	"fmt"
)

type Branch struct {
	ID          string `json:"Id"`
	Code        string `json:"Code"`
	Name        string `json:"Name"`
	IsBaseDepot bool   `json:"IsBaseDepot"`
	Inactive    bool   `json:"Inactive"`
}

func (a *Api) ListBranches(ctx context.Context) (*Response[[]*Branch], error) {
	tk, err := a.auth.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	res, err := get[responseCommon[[]*Branch]](
		ctx,
		fmt.Sprintf(`%s/api/v1/branchs/all`, a.config.url),
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

	return &Response[[]*Branch]{
		Code:    res.Code,
		Data:    res.Data,
		Total:   res.Total,
		Success: res.Success,
	}, nil
}
