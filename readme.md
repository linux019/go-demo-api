# Tasks

[![N|Solid](https://cldup.com/dTxpPi9lDf.thumb.png)](https://nodesource.com/products/nsolid)

Create a new Go application that does the following:
- Runs as a daemon on the command line- Structure the program so that it can be expanded for more routes and become a larger application, even though we're only going to put two routes in it for this test.
- Uses MongoDB as a database with an accounts collection
- Exposes port 80 and provides an HTTP REST interface
- Provides a /create_account route that takes an HTTP  POST with JSON body that contains an email address and password and creates a document in the accounts collection of the DB with the email and password
- Ensure the password is stored securely in the database
- Provides an /authenticate route that takes an HTTP POST with a JSON body that contains an email address and password and authenticates it vs a stored account.  Returns an HTTP 200 if successful or an HTTP 401 if not successful.
- Should handle errors and return 400's
- Should log nicely to the console
- Should have tests

# Usage
```
  -p int
        Server port (default 80)
  -password string
        Mongo Database password (default "root")
  -user string
        Mongo Database user (default "root")
```

# Notes
80 port is priveleged port and can be opened by root. So by default use another port

# Tests
Run a server
```
go test
```
If server uses another port, use *-p* argument to define the port
```
$ go test -v -args -p 4488
```
