# Requirements manager

A way of grouping of and tracking requirements between systems

- run the project go run main.go
- testing go test ./...

## TODO

- [ ] APIs
    - [ ] requirements
        - [ ] get all paged
    - [ ] products
        - [ ] get all paged
    - [ ] implementations
        - [ ] get all paged
        - [ ] put
        - [ ] docs

- [ ] refactor to standard go project

- [ ] front end
    - [ ] requirements
        - [ ] list
        - [ ] create
        - [ ] update
    - [ ] products list
    - [ ] implementation list

- [ ] tests
- [ ] make
- [ ] auth



### Some day

* swagger docs to use generated DTOs instead of the ent structs directly


### Swagger docs

- Accessed via http://localhost:8080/swagger/index.html
- https://goswagger.io/go-swagger/reference/annotations/params/
- https://github.com/swaggo/swag
- go install github.com/swaggo/swag/cmd/swag@latest
- go get -u github.com/swaggo/echo-swagger
- go get -u github.com/swaggo/swag // otherwise error unknown field LeftDelim in struct literal of type "github.com/swaggo/swag".Spec

- swag fmt // formats the annotations
- swag init // generates docs

### Ent files

- in ent/schema add a file for the entity to create
- run go generate ./ent

initially ent generate ./ent/schema before I had any files