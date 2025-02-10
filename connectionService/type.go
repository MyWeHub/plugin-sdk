package connectionService

import "context"

type IConnectionService interface {
	GetConnection(ctx context.Context, id string) (*ConnectionMessage, error)
	GetConnectionWithJWT(ctx context.Context, id string) (*ConnectionMessage, error)
	Close() error
}
