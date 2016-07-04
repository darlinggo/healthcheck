package healthcheck

import (
	"net/http"

	"golang.org/x/net/context"
)

type Checks struct {
	Checks  []Checker
	Context context.Context
	Logger  func(format string, msg ...interface{})
}

func NewChecks(ctx context.Context, logger func(format string, msg ...interface{}), checks ...Checker) Checks {
	return Checks{
		Checks:  checks,
		Context: ctx,
		Logger:  logger,
	}
}

func (c Checks) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, check := range c.Checks {
		err := check.Check(c.Context)
		if err != nil {
			if c.Logger != nil {
				c.Logger("Error performing health check for %T (%s): %+v\n", check, check.LogInfo(c.Context), err)
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Everything is on fire and nothing is okay."))
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
