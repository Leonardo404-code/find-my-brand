package brand

import "context"

type Services interface {
	SearchTerms(ctx context.Context, terms []byte) (*Result, error)
}
