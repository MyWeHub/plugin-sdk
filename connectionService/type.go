package connectionService

import "context"

type IConnectionService interface {
	GetConnection(ctx context.Context, id string) (*connectionMessage, error)
	GetConnectionWithJWT(ctx context.Context, id string) (*connectionMessage, error)
	Close() error
}
