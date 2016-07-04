package healthcheck

import "golang.org/x/net/context"

type Checker interface {
	Check(ctx context.Context) error
	LogInfo(ctx context.Context) string
}
