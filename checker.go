package healthcheck

import "golang.org/x/net/context"

// Checker defines the methods necessary to determine if a
// certain aspect of a service should be considered healthy.
// If the Check function returns nil, the instance is
// healthy. Otherwise, it's unhealthy. The LogInfo method
// should return a string that contains debug information
// or other useful identifying information for the log
// output when Check returns an error.
type Checker interface {
	Check(ctx context.Context) error
	LogInfo(ctx context.Context) string
}
