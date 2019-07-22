## High level
This application exposes functionality to create, read, update and delete a customer. 
At a high level, the application is composed of workflows (see /internal/workflow). Each 
workflow is a vertically complete customer feature. Workflows are pipelines that compose even smaller bits of functionality for example storage, event publishing etc. The entire set of workflows in the application are encapsulated in an 'app.Customer' (see internal/app). 'This encapsulation allows us to create different versions of the app. For instance, we are able to create an 'inmemapp' (see 'internal/app/inmem.go') which uses an inmemory storage mechanism and hence usueful during testing. The application is then wired up in a main function such as in '/cmd/customer-svc/main.go'.

<img src="./architecture.svg?sanitize=true">

### Folder organisation
- cmd: contains the application executables
- config: contains application level configuration files such as configurations for mongo, server ports etc. 
- internal: contains all the pieces of code internal to this application. Code here will not be shared with other applications. 
- postman: contains postman collections and environent variables to facilitate interactions with the applications APIs
- script: contains utility scripts for the application. e.g for managing deployments
- test: contains integration tests for the application.
