package cukcuk

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func post[RES any](ctx context.Context, url string, heades map[string]string, body any) (*RES, error) {
	var req *http.Request

	if body != nil {
		buf := new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}

		r, err := http.NewRequestWithContext(ctx, http.MethodPost, url, buf)
		if err != nil {
			return nil, err
		}
		req = r
	} else {
		r, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
		if err != nil {
			return nil, err
		}
		req = r
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
