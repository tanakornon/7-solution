package port

import "context"

type BaconIpsumAPI interface {
	GetText(ctx context.Context) (string, error)
}
