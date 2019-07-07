# customer-app
A simple CRUD application to aid in learning go


### Running the app
``` go install github.com/albertowusuasare/customer-app/... && customer-svc ```

### Build and run docker container
``` docker build -t customer-svc . && docker run -d -p 5090:5090 customer-svc```

### Running all tests
``` go test -v github.com/albertowusuasare/customer-app/... ```

### Running integration tests
```go test -v github.com/albertowusuasare/customer-app/cmd/customer-svc/test/integration```

### Retrieving go report card
``` https://goreportcard.com/report/github.com/albertowusuasare/customer-app```