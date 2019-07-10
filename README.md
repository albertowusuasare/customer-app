# customer-app
A simple CRUD application to aid in learning go

### Running the app
``` go install github.com/albertowusuasare/customer-app/... && customer-svc ```

### Build docker image
``` docker build --rm -f "Dockerfile" -t customer-app:0.0.x .```

### Run docker app
``` docker run -d  -p5090:5090 customer-app:0.0.x```

### Running all tests
``` go test -v github.com/albertowusuasare/customer-app/... ```

### Running integration tests
```go test -v github.com/albertowusuasare/customer-app/cmd/customer-svc/test/integration```

### Retrieving go report card
``` https://goreportcard.com/report/github.com/albertowusuasare/customer-app```

### Running pre commit sanity check script
``` ./script/pre_commit.sh ```

