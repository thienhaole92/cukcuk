package token

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Token struct {
	config *Config
}

func New(config *Config) *Token {
	return &Token{
		config: config,
	}
}

type Response struct {
	Code int `json:"Code"`
	Data struct {
		AccessToken string `json:"AccessToken"`
		AppID       string `json:"AppID"`
		CompanyCode string `json:"CompanyCode"`
		Domain      string `json:"Domain"`
	} `json:"Data"`
	Success bool `json:"Success"`
	Total   int  `json:"Total"`
}

func (t *Token) GetToken(ctx context.Context) (string, error) {
	signature, loginTime, err := t.buildSignature()
	if err != nil {
		return "", err
	}

	type Body struct {
		AppID         string `json:"AppID"`
		Domain        string `json:"Domain"`
		LoginTime     string `json:"LoginTime"`
		SignatureInfo string `json:"SignatureInfo"`
	}

	res, err := t.post(
		ctx,
		fmt.Sprintf(`%s/api/account/login`, t.config.url),
		map[string]string{},
		Body{
			AppID:         t.config.appID,
			Domain:        t.config.domain,
			LoginTime:     loginTime,
			SignatureInfo: signature,
		},
	)

	if err != nil {
		return "", err
	}

	return res.Data.AccessToken, nil
}

func (t *Token) buildSignature() (string, string, error) {
	loginTime := time.Now().String()

	msgObj := struct {
		AppID     string `json:"AppID"`
		Domain    string `json:"Domain"`
		LoginTime string `json:"LoginTime"`
	}{
		AppID:     t.config.appID,
		Domain:    t.config.domain,
		LoginTime: loginTime,
	}

	message, err := json.Marshal(msgObj)
	if err != nil {
		return "", "", err
	}

	key := []byte(t.config.secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	signature := hex.EncodeToString(h.Sum(nil))

	return signature, loginTime, nil
}

func (t *Token) post(ctx context.Context, url string, heades map[string]string, body any) (*Response, error) {
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

		var result Response
		if err := json.Unmarshal(body, &result); err != nil {
			return nil, err
		}

		return &result, nil
	}
}
