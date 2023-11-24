package cukcuk

import "context"

type TokenClient interface {
	GetToken(ctx context.Context) (string, error)
}
