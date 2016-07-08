package healthcheck

import (
	"net/http"

	"golang.org/x/net/context"
)

// Checks defines a group of Checkers, a log function to
// write their errors with, and a context to associate
// with the Checkers. Checks is an http.Handler.
type Checks struct {
	Checks []Checker
	// This property will be deprecated once Go 1.7 is released
	Context context.Context
	Logger  func(format string, msg ...interface{})
}

// NewChecks returns a Checks instance using the passed
// context, logging function, and Checkers.
func NewChecks(ctx context.Context, logger func(format string, msg ...interface{}), checks ...Checker) Checks {
	return Checks{
		Checks:  checks,
		Context: ctx,
		Logger:  logger,
	}
}

// ServeHTTP fulfills the http.Handler interface. Calling
// ServeHTTP will call the Check method on all the
// Checkers associated with the Checks. If any of the
// Checkers returns an error, a response with status 500
// is written and the error is logged. Otherwise, a response
// with status 200 is written, and the text "OK" is written
// to the response.
func (c Checks) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var errFound bool
	for _, check := range c.Checks {
		err := check.Check(c.Context)
		if err != nil {
			if c.Logger != nil {
				c.Logger("Error performing health check for %T (%s): %+v\n", check, check.LogInfo(c.Context), err)
			}
			errFound = true
		}
	}
	if errFound {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Everything is on fire and nothing is okay."))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
