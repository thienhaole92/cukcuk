# CukCuk Go

A Go wrapper for [CukCuk APIs](https://graphapi.cukcuk.vn/document/api/index.html)

You need to get an API key from [CukCuk](https:/cukcuk.vn)

## Usage

### Add the module to your go project

```sh
go get github.com/thienhaole92/cukcuk/cukcuk
```

### Sample Code

### Get access token

```go
package main

import (
	"context"
	"github.com/thienhaole92/cukcuk/token"
	"go.uber.org/zap"
)

const CUKCUK_URL = "https://graphapi.cukcuk.vn"
const CUKCUK_APP_ID = "CUKCUKOpenPlatform"
const CUKCUK_DOMAIN = "CUKCUK_DOMAIN"
const CUKCUK_SECRET = "CUKCUK_SECRET"

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()

	tkc := token.New(token.NewConfig(
		CUKCUK_URL,
		CUKCUK_APP_ID,
		CUKCUK_DOMAIN,
		CUKCUK_SECRET,
	))
	token, _ := tkc.GetToken(context.Background())
	sugar.Infow("token obtained", "token", token)
}
```

### APIs

```go
package main

import (
	"context"

	"github.com/thienhaole92/cukcuk/cukcuk"
	"go.uber.org/zap"
)

type TestTokenClient struct {
	token string
}

func NewTestTokenClient(token string) *TestTokenClient {
	return &TestTokenClient{
		token: token,
	}
}

func (t *TestTokenClient) GetToken(ctx context.Context) (string, error) {
	return t.token, nil
}

const ACCESS_TOKEN = "ACCESS_TOKEN"
const COMPANY_CODE = "COMPANY_CODE"

func main() {
	tk := NewTestTokenClient(ACCESS_TOKEN)
	config := cukcuk.NewConfig("https://graphapi.cukcuk.vn", COMPANY_CODE)
	cli := cukcuk.New(config, tk)

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()

	res, err := cli.SAInvoicePaging(context.Background(), &cukcuk.SainvoicePagingReq{
		Page:         1,
		Limit:        100,
		HaveCustomer: true,
	})
	if err != nil {
		sugar.Error(err)
	} else {
		sugar.Infow("response", "data", res.Data)
	}
}
```
