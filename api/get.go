package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func get[RES any](ctx context.Context, url string, heades map[string]string) (*RES, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range heades {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 400 {
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var httpErr error
		if err := json.Unmarshal(body, &httpErr); err != nil {
			return nil, err
		}

		return nil, httpErr
	} else {
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var result RES
		if err := json.Unmarshal(body, &result); err != nil {
			return nil, err
		}

		return &result, nil
	}
}
