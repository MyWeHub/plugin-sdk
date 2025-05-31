package offline

import "google.golang.org/protobuf/proto"

type IOffline interface {
	Serve() error
	SetConfig(c proto.Message)
}
