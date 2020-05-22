# Notes while learning mux framework

## TODO

- Read net/http package
- Read mux API
- Middleware Implementations
- Testing http packages
- Firestore API

## Notes
- Use code organization just like being done in the project

### Clean Architecture Fundamentals
- Independent of frameworks
- Testable
- Independent of UI
- Independent of Database
  - Keep an interface aside and implement acc to the db specs
  - Example, firestore-repo can have its own implementation of interface
    and mysql-repo can have another implementation
  - Now you can pass these handlers(any database) to the module where it is used
    (Dependency Injection). It will become independent of the type of db used.
  - A general process is :
    - Step 1 - create an interface
    - Step 2 - implement that interface with your service of choice
    - Step 3 - pass the sevice in the modules where it is used
- Independent of any external agency