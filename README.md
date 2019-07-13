# customer-app
A simple CRUD application to aid in learning go

### Running the app
``` make install run```

### Build docker image
``` docker build --rm -f "Dockerfile" -t customer-app:0.0.x .```

### Run docker app
``` docker run -d  -p5090:5090 customer-app:0.0.x```

### Running all tests
``` make test```

### Running integration tests
```make int-test```

### Retrieving go report card
``` https://goreportcard.com/report/github.com/albertowusuasare/customer-app```

### Running pre commit sanity check script
``` make sanity-check```

