# healthcheck

healthcheck is a package that supplies an `http.Handler` that reports whether your service is healthy or not. All you have to do is define a function that returns an error if one of your service's dependencies is unhealthy.

## Using the package

### Defining a check

To use the package, first we need to know if our service is healthy or not. A service is healthy if all the checks for it return no error. Let's define a new check:

```go
type SQL struct {
	DB *sql.DB
	ID string
}

func (s SQL) Check(ctx context.Context) error {
	return s.DB.Ping()
}

func (s SQL) LogInfo(ctx context.Context) string {
	return s.ID
}
```

This check fulfills the `healthcheck.Checker` interface, by defining the `Check` method and the `LogInfo` method. The `Check` method returns nil if the check should be considered successful, or an error if it should be marked unhealthy. The `LogInfo` method returns information about the instance of the check (a database host, a connection string, or what have you) to uniquely identify the check in logs.

You can use the SQL check defined above by calling `NewSQL` from this package.

### Setting up the `http.Handler`

Now that we have a check, we want an `http.Handler` that will return a 500 status code if the check fails, so we know something is wrong and our monitoring can alert.

For that, we use the `NewChecks` function:

```go
// db is set up as an *sql.DB connection
sqlCheck := SQL{DB: db, ID: "mydb"}
handler := NewChecks(context.Background(), nil, sqlCheck)
```

`handler` is now an `http.Handler`, and will return a 500 status code if `sqlCheck` returns an error. You can set up as many checks on the handler as you like:

```go
// db is set up as an *sql.DB connection
// otherdb is set up as a separate connection
sqlCheck := SQL{DB: db, ID: "mydb"}
otherCheck := SQL{DB: otherdb, ID: "myotherdb"}
handler := NewChecks(context.Background(), nil, sqlCheck, otherCheck)
```

Of course, it's not super useful to know that your service is unhealthy, but not know what the error was. To get that info, we need to pass a logging function. A logging function is anything that fills the `func(message string, ...args interface{})` signature. You'll notice that `log.Printf` fits the bill. Let's use that.

```go
// db is set up as an *sql.DB connection
// otherdb is set up as a separate connection
sqlCheck := SQL{DB: db, ID: "mydb"}
otherCheck := SQL{DB: otherdb, ID: "myotherdb"}
handler := NewChecks(context.Background(), log.Printf, sqlCheck, otherCheck)
```

Now the error that was returned that caused the 500 will be logged.
